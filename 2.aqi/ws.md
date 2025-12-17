# ws
websocket 连接是基于tcp的http协议的升级协议，可以在客户端和服务端之间建立长连接，可以持续、双向的进行通信。在aqi框架中，
需要首先使用http请求升级ws连接， 在使用action字段进行路由访问，形式为：a.b.c，使用json作为通信载体。
对后端而言，ws.HttpHandler(c.Writer, c.Request)升级http为ws，使用ws.NewRouter()构建路由实例，使用Add,Group,Use
三个函数进行路由的延伸、分组以及中间件的使用。

## ws.NewRouter
使用map管理路由（根据客户端发送不同的action，执行不同的handlerFunc）

## add
新增路由，就像http的url一样，设置了action的值，根据输入匹配路由然后返回
响应

## group
使用[]string管理路由，在add的时候就会把前面所有的路由通过.的方式连接形成
链式路由，例如group1.group2.hi。

## use
使用[]handlerfunc管理中间件，使用use给分组路由增加中间件，使用a.next执行下一个
handlerfunc，a.Abort终止执行。（和gin框架的中间件注册和使用类似）