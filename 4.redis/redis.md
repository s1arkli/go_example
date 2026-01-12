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

    if err := cn.WithReader(c.context(ctx), c.cmdTimeout(cmd), func(rd *proto.Reader) error {
		...
        return readReplyFunc(rd)
    })

    func (cmd *StringCmd) readReply(rd *proto.Reader) (err error) {
        cmd.val, err = rd.ReadString()
		return err
    }
```
- writeCmd(wr, cmd)的作用是将cmd的args以RESP的规范序列化写入wr待用。其中cn是封装了tcp连接的*pool.Conn，cn.WithWriter的作用是将完成序列化
的内容写入缓冲区，调用Flush将缓存区的内容经tcp发送给redis，redis执行命令，实现crud。
- WithReader是从redis拿本次cmd执行和响应，实际执行函数是readReply，返回的数据也是符合RESP规范的，需要反序列化到cmd.val字段内部，通过
cmd.Result拿到响应结果。

# 总结
- 此包就是把符合RESP协议的redis命令，高度抽象为go函数，每次调用go函数，就是将输入的命令和参数编码为redis命令，并将命令经tcp连接传输给redis。

- **redis中使用map/struct**
  - 关于结构体和map的使用，结构体是具有固定字段，在序列化为json的时候需要定义结构体，并且需要json tag。而map没有固定字段，以k-v的形式，灵活性
  更好。具体使用哪一个应该是看需要解析的内容的结构，如果是固定结构或者变动不频繁，例如mysql表的model，就可以使用struct形式。如果保存的是需要经
  常拓展或者结构变动频繁的内容，就适合map，临时定义struct不现实。还有就是使用map读取缓存的时候，需要大量的类型断言，而使用struct可以避免这样的情况。
  - 在meet项目中，由于我们对于某些model的信息需求频繁，所以将model(struct)序列化为json保存进redis中，在get的时候先从缓存拿，再反序列化为定义
  好的model。
  - json序列化的必要性：Redis保存的是字节，而json可以直接被计算机转换为字节，在存入Redis时不会进行其他转换，保证存进去的是什么，取出来的就是什么。

## tips
- 把结构体的方法赋值给其他结构体的函数字段，这样结构体就可以通过字段间接调用方法（这个字段就像是一个接口一样，满足这样的函数形式的函数都可以赋值给这个
结构体，这个结构体就可以调用不同的方法）。

- 类型断言判断接口是否实现
```go
var _ Pooler = (*ConnPool)(nil)
```
var一个Pooler类型变量，将nil类型断言为*ConnPool类型，可以在编译之前判断定义的ConnPool是否实现了Pooler接口。

- 连接池：
    - 配置opt--连接超时、最小最大连接数量、连接的持续时间等等，所有的配置都从这个字段读取
    - 读写缓存区大小--过大缓存长时间在缓存区，造成读写延迟高。过小导致频繁调用缓存区，导致资源浪费。
    - 错误保存
    - 并发锁--很可能高频使用连接，需要增加锁来保证并发安全。
    - 连接数组--保存所有的tcp连接，方便进行关闭和新增连接
    - 钩子函数--方便在连接前后增加自己需要的逻辑，如在连接前后增加日志

