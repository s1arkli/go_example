# redis
go-redis包的主要作用是把函数语言序列化为符合RESP规范的redis命令，然后通过tcp连接发送给redis，实现crud。
例如get key序列化如下
*2\r\n
$3\r\n
get\r\n
$3\r\n
key\r\n

# Client
```go
type Client struct {
	//
	*baseClient
	//process func
	cmdable
}
```
Client结构体封装了连接池（用于管理tcp连接）、对客户端的配置（超时等等）、读写锁...，其中cmdable字段是函数类型（这里在init的时候
c.cmdable = c.Process被Process赋值，后续会使用到）

## Set
把set+key+value+ttl组装成一个数组然后存到StatusCmd结构体中，执行client.Process(也就是在init时赋值给cmdable字段的函数)

## Get
把get+key组装起来，执行process

## Del 
把del+key组装起来，执行process

## _process
最终将命令写入redis的函数

- 重点部分
```go
    if err := cn.WithWriter(c.context(ctx), c.opt.WriteTimeout, func(wr *proto.Writer) error {
        return writeCmd(wr, cmd)
  }); err != nil {...}

func (cn *Conn) WithWriter(ctx context.Context, timeout time.Duration, fn func(wr *proto.Writer) error,) error {
    ...
	//marshal cmd
    if err := fn(cn.wr); err != nil {
        return err
    }
    //send cmd by tcp
    return cn.bw.Flush()
}
```
writeCmd(wr, cmd)的作用是将cmd的args以RESP的规范序列化写入wr待用。其中cn是封装了tcp连接的*pool.Conn，cn.WithWriter的作用是将完成序列化
的内容写入缓冲区，调用Flush将缓存区的内容经tcp发送给redis，redis执行命令，实现crud。

# 总结
- 此包就是把redis能看懂的符合RESP协议的命令，高度抽象为go函数，每次调用go函数，就是将输入的命令和参数编码为redis命令，并将命令经tcp连接传输
给redis。

- 关于结构体和map的使用，结构体是具有固定字段，在序列化为json的时候需要有结构体实例，并且需要json tag。而map没有固定字段，以k-v的形式，
可以保存任意k-v信息，灵活性更好。我觉得应该是看需要解析的内容的结构，如果是固定结构，例如mysql表内的固定struct，就可以使用struct形式。如果保存
的是需要经常拓展或者修改的内容，就适合map，临时定义struct不现实。在meet项目中，我们将最常用的user结构体存入redis，在get的时候也是用user结构体
去解析，这样就省下读取mysql的步骤，直接从redis内部读取。

- 连接池：
    - 配置opt--连接超时、最小最大连接数量、连接的持续时间等等，所有的配置都从这个字段读取
    - 读写缓存区大小--过大缓存长时间在缓存区，造成读写延迟高。过小导致频繁调用缓存区，导致资源浪费。
    - 错误保存
    - 并发锁--很可能高频使用连接，需要增加锁来保证并发安全。
    - 连接数组--保存所有的tcp连接，方便进行关闭和新增连接
    - 钩子函数--方便在连接前后增加自己需要的逻辑，如在连接前后增加日志

- 学到了把结构体的方法赋值给其他结构体的函数字段，这样两个函数都可以调用这个方法（这个字段就像是一个接口一样，满足这样的函数形式的函数都可以赋值给这个
结构体，这个结构体就可以调用不同的方法）。

- 类型断言判断接口是否实现
```go
var _ Pooler = (*ConnPool)(nil)
```
var一个Pooler类型变量，将nil类型断言为*ConnPool类型，可以在编译之前判断定义的ConnPool是否实现了Pooler接口。