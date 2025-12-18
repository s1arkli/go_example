# ws
websocket 连接是基于tcp的http协议的升级协议，可以在客户端和服务端之间建立长连接，可以持续、双向的进行通信。

在aqi框架中，需要首先使用http请求升级ws连接， 在使用action字段进行路由访问，形式为：a.b.c，使用json作为通信载体。
对后端而言，ws.HttpHandler(c.Writer, c.Request)升级http为ws，使用ws.NewRouter()构建路由实例，使用Add,Group,Use
三个函数进行路由的延伸、分组以及中间件的使用。

## ws.NewRouter

```go
type Routers struct {
    manager        *ActionManager
    handlerMembers HandlersChain
    groups         []string
}
```
返回的routers实例---manager是map，形式为name:[]handlerFunc的键值对（一是管理当前路由下的所有handlerFunc，二是由于map的
唯一键可以防止重复注册路由）---handlerMembers也是[]handlerFunc，作用是使用use管理中间件。---groups保存的是当前路由（
由于Group方法是值类型接收者，每次都会复制一份Routers，所以内部保存的是父路由以及当前路由，也就是这个分组的路由）

## Add
新增子路由，先将groups和输入name拼出字路由的name，在manager中检查是否重复注册字路由。然后初始化[]handlerFunc，初始长度为use的
中间件个数，容量为中间件个数以及传入fn的数量(这里就定下了执行顺序，优先中间件，然后是add中输入的fn)。将中间件copy到chains中（这里
为什么要用copy，我的理解：add是创建子路由，而handlerMembers内部是整个group的中间件，属于共享资源，直接使用copy来复制一份副本可以
保证中间件数据的独立性。我觉得append(chains, r.handlerMembers...)也可以？）最后将fn append到chains中（保证了中间件和fn的执行
顺序），与name一起传入manager中形成键值对，完成子路由的创建，每当访问到这个路由，就会根据manage内的[]handleFunc顺序执行。

总结：Add是创建ws子路由的值类型方法，由于manage字段是指针类型，每次add都会修改父本（路由组？），由父本管理分组下的所有子路由。
## Group
值类型方法，每次会创建副本，相当于子路由组实例。返回的是接口类型（高度封装？只暴露方法而不是结构体，让使用者只专注于路由和函数，而不是
具体字段）

总结：Group是分组管理的必要方法，每次返回子路由实例，保证代码中逻辑清晰，路由可读性好。
## Use
值类型方法（多次use产生的实例进行Add会怎样？）返回一个append了中间件的接口实例，方便后续进行Add。

总结：Use是路由组使用中间件的方法，方便给整个路由组增加公共的前置的函数，例如：身份认证。
## aqi
aqi的ws服务，可以快速建立ws连接，进行前后端的数据交互。