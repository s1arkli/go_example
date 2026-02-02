## Dial
```go
func Dial(url string) (*Connection, error) {
	return DialConfig(url, Config{
		Locale: defaultLocale,
	})
}
```

根据配置的url，建立tcp连接。

```go
   amqp091.Connection
```
内部是封装了tcp连接、心跳机制（由于tcp突然断联的另一方是不会有察觉的，所以需要心跳机制，固定间隔时间发送心跳检测，保证连接的连通性），
、消息管道、并发控制（多个chan公用一个tcp，由于chan的开销低，tcp需要3次握手，4次挥手，所以多个chan共享一个tcp是最性能的设计），

```go
func (ch *Channel) QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args Table) (Queue, error) 
```
函数是向rabbitmq发送声明队列，根据是否需要等待相应区分返回值。（默认关闭nowait，也就是每次都是需要阻塞等待rabbitmq的响应）。在默认设置下，
返回队列名称，积压消息数量、当前订阅的消费者数量，配合k8s进行资源的扩容以及缩减，既保证不过度浪费系统资源，也保证在峰值请求下的正常运行。

```go
func (ch *Channel) PublishWithDeferredConfirm(exchange, key string, mandatory, immediate bool, msg Publishing) (*DeferredConfirmation, error)
```
sync.Mutex给区域加锁，保证多个goroutine同时只有一个能够访问conn，这也是多个chan共用一个tcp连接的关键，同一时刻只有一个chan能够向rabbitmq
发送符合AMQP协议的命令，实现声明队列，给指定队列发送消息等