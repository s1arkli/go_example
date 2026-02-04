# rabbitmq
这是一个订阅发布模型，消费者阻塞等待消息，发布者触发后发布消息，由此完成消息的传递。消息队列在实际生产中，可以作为微服务异步通信的桥梁。
**可以同时开启多个消费者和发布者，在demo中就是开启多个终端，消息的顺序不做保证。**

---

## Dial
```go
func Dial(url string) (*Connection, error) {
	return DialConfig(url, Config{
		Locale: defaultLocale,
	})
}
```
根据配置的url（内部包含rabbitmq端口，初始账号密码），建立tcp连接。

---

```go
   type Connection struct {
    destructor sync.Once  // shutdown once
    sendM      sync.Mutex // conn writer mutex
    m          sync.Mutex // struct field mutex

    conn io.ReadWriteCloser

    rpc       chan message
    writer    *writer
    sends     chan time.Time     // timestamps of each frame sent
    deadlines chan readDeadliner // heartbeater updates read deadlines

    allocator *allocator // id generator valid after openTune
    channels  map[uint16]*Channel

    noNotify bool // true when we will never notify again
    closes   []chan *Error
    blocks   []chan Blocking

    errors chan *Error
    // if connection is closed should close this chan
    close chan struct{}

    Config Config // The negotiated Config after connection.open

    Major      int      // Server's major version
    Minor      int      // Server's minor version
    Properties Table    // Server properties
    Locales    []string // Server locales

    closed int32 // Will be 1 if the connection is closed, 0 otherwise. Should only be accessed as atomic
}
```
内部是封装了tcp连接、心跳机制（由于tcp突然断联的另一方是不会有察觉的，所以需要心跳机制，固定间隔时间发送心跳检测，保证连接的连通性），
、消息管道、并发控制（多个chan公用一个tcp，由于chan的开销低，tcp需要3次握手，4次挥手，所以多个chan共享一个tcp是最性能的设计）

---

```go
func (c *Connection) Channel() (*Channel, error)

type Channel struct {
    destructor sync.Once
    m          sync.Mutex // struct field mutex
    confirmM   sync.Mutex // publisher confirms state mutex
    notifyM    sync.RWMutex

    connection *Connection

    rpc       chan message
    consumers *consumers

    id uint16

    // closed is set to 1 when the channel has been closed - see Channel.send()
    closed int32
    close  chan struct{}

    // true when we will never notify again
    noNotify bool

    // Channel and Connection exceptions will be broadcast on these listeners.
    closes []chan *Error

    // Listeners for active=true flow control.  When true is sent to a listener,
    // publishing should pause until false is sent to listeners.
    flows []chan bool

    // Listeners for returned publishings for unroutable messages on mandatory
    // publishings or undeliverable messages on immediate publishings.
    returns []chan Return

    // Listeners for when the server notifies the client that
    // a consumer has been cancelled.
    cancels []chan string

    // Allocated when in confirm mode in order to track publish counter and order confirms
    confirms   *confirms
    confirming bool

    // Selects on any errors from shutdown during RPC
    errors chan *Error

    // State machine that manages frame order, must only be mutated by the connection
    recv func(*Channel, frame)

    // Current state for frame re-assembly, only mutated from recv
    message messageWithContent
    header  *headerFrame
    body    []byte
}
```
- channel内部包含互斥锁、tcp连接、消息管道、绑定的消费者、各种通知管道（比如错误通知管道，用户可以监听此管道，对错误进行处理）
- channel可以调用声明队列的方法，也可以绑定消费者
- 一个tcp连接可以复用多个channel，即使channel发生错误，只影响当前的channel。由于多个tcp连接的资源占用大，所以使用一个tcp复用多个channel，
这样保证了资源的最大化利用，以及服务的稳定性。

---

```go
func (ch *Channel) QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args Table) (Queue, error) 
```
函数是向rabbitmq发送声明队列，根据是否需要等待相应区分返回值。（默认关闭nowait，也就是每次都是需要阻塞等待rabbitmq的响应）。在默认设置下，
返回队列名称，积压消息数量、当前订阅的消费者数量，配合k8s进行资源的扩容以及缩减，既保证不过度浪费系统资源，也保证在峰值请求下的正常运行。

---

```go
func (ch *Channel) PublishWithDeferredConfirm(exchange, key string, mandatory, immediate bool, msg Publishing) (*DeferredConfirmation, error)
```
sync.Mutex给区域加锁，保证多个goroutine同时只有一个能够访问conn，这也是多个chan共用一个tcp连接的关键，同一时刻只有一个chan能够向rabbitmq
发送符合AMQP协议的命令，实现声明队列，给指定队列发送消息等

```go
func (ch *Channel) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args Table) (<-chan Delivery, error)

func (subs *consumers) buffer(in chan *Delivery, out chan Delivery) {
    defer close(out)
    defer subs.Done()

    var inflight = in
    var queue []*Delivery
    for delivery := range in {
        queue = append(queue, delivery)

	    for len(queue) > 0 {
            select {
            case <-subs.closed:
            return

            case delivery, consuming := <-inflight:
            if consuming {
            queue = append(queue, delivery)
            } else {
            inflight = nil
            }

            case out <- *queue[0]:
            queue[0] = nil
            queue = queue[1:]
	        }
        }
}
```
消费者获取消息的流程。buffer是获取消息的主要函数，in是直接接受消息的管道，out是用户获取消息的管道，这里做了一步缓冲，大概的作用应该是避免前者
出问题直接影响获取，使得消费者-->缓冲区-->rabbitmq这之间尽量保持相对独立。

这里的双循环很有意思，外层遍历in管道，在in管道接收到数据时进入循环，内层循环被触发之后，会不断从in拿取消息并将消息塞给out（也就是消费者），内层
循环终止随即进入下一轮的等待。

var了一个inflight，当in管道出现问题将inflight赋值为nil终止循环，不直接操作in，并安全的终止循环。
