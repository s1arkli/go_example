

## ws.NewRouter
使用map管理路由（根据客户端发送不同的action，执行不同的handlerFunc）

## add
新增路由，就像http的url一样，设置了action的值，根据输入匹配路由然后返回
响应

## group
使用[]string管理路由，在add的时候就会把前面所有的路由通过.的方式连接形成
链式路由，例如group1.group2.hi。

## use
在Routers结构体中对中间件做管理，按照传入的顺序执行。