## ws概述
- websocket 连接是基于tcp的http协议的升级协议，可以在客户端和服务端之间建立长连接，可以持续、双向的进行通信。

## aqi中的ws概述
- 在aqi框架中，需要首先使用http请求升级ws连接， 在使用action字段进行路由访问，形式为：a.b.c，使用json作为通信载体。
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
- 返回的routers实例---manager是map，形式为name:[]handlerFunc的键值对（一是管理当前路由下的所有handlerFunc，二是由于map的
唯一键可以防止重复注册路由）---handlerMembers也是[]handlerFunc，作用是使用use管理中间件。---groups保存的是当前路由（
由于Group方法是值类型接收者，每次都会复制一份Routers，所以内部保存的是父路由以及当前路由，也就是这个分组的路由）

## Add
- 新增子路由，先将groups和输入name拼出字路由的name，在manager中检查是否重复注册字路由。然后初始化[]handlerFunc，初始长度为use的
中间件个数，容量为中间件个数以及传入fn的数量(这里就定下了执行顺序，优先中间件，然后是add中输入的fn)。将中间件copy到chains中（这里
为什么要用copy，我的理解：add是创建子路由，而handlerMembers内部是整个group的中间件，属于共享资源，直接使用copy来复制一份副本可以
保证中间件数据的独立性。我觉得append(chains, r.handlerMembers...)也可以？）最后将fn append到chains中（保证了中间件和fn的执行
顺序），与name一起传入manager中形成键值对，完成子路由的创建，每当访问到这个路由，就会根据manage内的[]handleFunc顺序执行。

- 总结：Add是创建ws子路由的值类型方法，由于manage字段是指针类型，每次add都会修改父本（路由组？），由父本管理分组下的所有子路由。
## Group
- 值类型方法，每次会创建副本，相当于子路由组实例。返回的是接口类型（高度封装？只暴露方法而不是结构体，让使用者只专注于路由和函数，而不是
具体字段）

- 总结：Group是分组管理的必要方法，每次返回子路由实例，保证代码中逻辑清晰，路由可读性好。
## Use
- 值类型方法（多次use产生的实例进行Add会怎样？）返回一个append了中间件的接口实例，方便后续进行Add。

总结：Use是路由组使用中间件的方法，方便给整个路由组增加公共的前置的函数，例如：身份认证。

# client结构体

```go 
type Client struct {
Hub            *Hubc       `json:"-"`
Conn           net.Conn    `json:"-"`
Send           chan []byte `json:"-"`
Endpoint       string      `json:"-"` //入口地址
OnceId         string      `json:"-"` //临时ID，扫码登录等场景作为客户端唯一标识
ClientId       string      `json:"-"` //客户端ID
Disconnecting  bool        `json:"-"` //已被设置为断开状态（消息发送完之后断开连接）
SyncMsg        bool        `json:"-"` //是否接收消息
LastMsgId      int         `json:"-"` //最后一条消息ID
RequiredValid  bool        `json:"-"` //人机验证标识
Validated      bool        `json:"-"` //是否已验证
ValidExpiry    time.Time   `json:"-"` //验证有效期
ValidCacheData any         `json:"-"` //验证相关缓存数据
AuthCode       string      `json:"-"` //用于校验JWT中的code，如果相等识别为同一个用户的网络地址变更
ErrorCount     int         `json:"-"` //错误次数
Closed         bool        `json:"-"` //是否已经关闭

Limiter      *rate.Limiter `json:"-"` //限速器
RequestQueue chan string   `json:"-"` //处理队列

HttpRequest *http.Request       `json:"-"`
HttpWriter  http.ResponseWriter `json:"-"`

User              *User     `json:"user,omitempty"`    //关联用户
Scope             string    `json:"scope"`             //登录jwt scope, 用于判断用户从哪里登录的
AppId             string    `json:"appId"`             //登录应用Id
StoreId           uint      `json:"storeId"`           //店铺ID
MerchantId        uint      `json:"merchantId"`        //商户ID
TenantId          uint      `json:"tenantId"`          //租户ID
Platform          string    `json:"platform"`          //登录平台
GroupId           string    `json:"groupId"`           //用户分组Id
IsLogin           bool      `json:"isLogin"`           //是否已登录
LoginAction       string    `json:"loginAction"`       //登录动作
ForceDialogId     string    `json:"forceDialogId"`     //打开聊天界面的会话ID
IpAddress         string    `json:"ipAddress"`         //IP地址
IpAddressPort     string    `json:"IpAddressPort"`     //IP地址和端口
IpLocation        string    `json:"ipLocation"`        //通过IP转换获得的地理位置
ConnectionTime    time.Time `json:"connectionTime"`    //连接时间
LastRequestTime   time.Time `json:"lastRequestTime"`   //最后请求时间
LastHeartbeatTime time.Time `json:"lastHeartbeatTime"` //最后发送心跳时间

mu   sync.RWMutex
Keys map[string]any

// recent logs ring buffer (last 100 items)
recentLogs  [100]string
recentIdx   int
recentCount int
}
```

客户端结构体主要记录了ws连接、所属hub、发送消息管道（方便发送响应）、客户端的各种状态和信息（由于ws是无状态的，服务器方需要维护状态，不然
有一边退出了ws连接，另一边仍然不知道，继续发送消息会导致panic或报错）、是否被注销等字段。

## Reader()
- wsutil.ReadClientData监听ws连接，获取客户端请求。对text类型的请求放到RequestQueue管道做缓冲。请求为ping返回pong（测试连通性）。
其余类型记录日志不做处理。

- 总结：Reader是读取请求的方法

## Request()
- 遍历RequestQueue管道，管道一有值就会处理请求。这里做了一个防止恶意的过量请求。正常的请求进入Dispatcher。中间穿插各种参数的校验以及日志的记录。
### Dispatcher(c,req)主要流程
- 首先将req序列化，拿到id，param，action。InitManager().Handlers(req.Action)（全局管理的manager用key=req.Action拿到handlersChain
，其中包括这个路由下的所有handlerFunc）.然后执行第一个handleFunc，接着调用c.Next，通过context内部的index字段与handleFuncs的长度进行比对，
每次调用执行index++，保证能够完全执行所有的handleFunc。

- 总结：Request是处理请求的主要方法，通过监听ws链接拿到初始请求，序列化req转成适合开发的结构体。最终将action字段与manager map进行匹配拿到
待执行函数，依次执行。

## Write()
- 等待send管道，一有数据(事先处理好的json内容)调用wsutil.WriteServerMessage将数据通过conn传输给客户端。这里还做了一个定时器，每5秒给客户端
发送ping，告诉客户端连接未中断，同时记录时间和日志（这里触发代表这个流程过长或者阻塞需要排查问题）。

- 总结：Write()是服务器向客户端发送消息的方法，同时记录发送时长，方便流程优化。

## ws.HttpHandler()
- 同时开启Reader，Write，Request三个协程，保证客户端和服务器在连通期间的交互与数据传输。

# Hubc

```go
type Hubc struct {
	//访客列表
	Guests []*Client

	//已登录用户 map[string]*User
	Users *sync.Map

	//用户数统计
	LoginCount int
	GuestCount int

	//发布订阅
	PubSub *PubSub

	//登录和断开通道
	Connection chan *Client
	Disconnect chan *Client
}
```

