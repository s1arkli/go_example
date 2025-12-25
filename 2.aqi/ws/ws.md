## aqi中的ws概述
- 对于已连接用户有两种接收消息的方式，一种是通过客户端访问指定路由，ws连接返回客户端需要的数据。一种是让用户订阅Topic，服务器推送消息。
  这两种方式的区别在于一个是被动响应，一个是主动推送，这也是ws的优点之一（双向通信）。
- 在aqi框架中，需要首先使用http请求升级ws连接， 在使用action字段进行路由访问，形式为：a.b.c，使用json作为通信载体。
对后端而言，ws.HttpHandler(c.Writer, c.Request)升级http为ws，使用ws.NewRouter()构建路由实例，使用Add,Group,Use
三个函数进行路由的延伸、分组以及中间件的使用。当客户端访问指定路由时，服务器做出响应（操作数据库，调用第三方，返回消息等）。
- 在aqi中，对于ws连接的管理有一个层级的关系，hubc是总枢纽，其下记录着users，而users保存对应的clients（多端登陆的conn），在hubc中又管理着
pubsub系统（是服务器主动推送消息的管理中枢，初始化hubc时会开启goroutine来推送消息）。当某个Topic需要发送消息的时候，hubc会找到订阅该Topic
下所有的user，并把消息发送给user的所有clients。

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

- **主要函数**

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
...

Limiter      *rate.Limiter `json:"-"` //限速器
RequestQueue chan string   `json:"-"` //处理队列

HttpRequest *http.Request       `json:"-"`
HttpWriter  http.ResponseWriter `json:"-"`

User              *User     `json:"user,omitempty"`    //关联用户
...
}
```

客户端结构体主要记录了ws连接、所属hub、发送消息管道（方便发送响应）、客户端的各种状态和信息（由于ws是无状态的，服务器方需要维护状态，不然
有一边退出了ws连接，另一边仍然不知道，继续发送消息会导致panic或报错）、是否被注销等字段。

- **主要函数**

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

var Hub *Hubc

func NewHubc() *Hubc {
        Hub = &Hubc{
        PubSub:     NewPubSub(),
        Guests:     []*Client{},
        Users:      new(sync.Map),
        Connection: make(chan *Client),
        Disconnect: make(chan *Client),
    }
    return Hub
}
```

使用var初始化hubc结构体，使用NewHubc给hub进行初始化配置，管理user连接和断联、topic消息发布中枢。

- **主要函数**

## Run()
- ```go go h.PubSub.Start() ```
开启goroutine用来监听订阅，一旦开启了订阅，则会对订阅进行处理，发送消息给用户

- ```go go h.guard()```
每30秒检测一次，如果守护不为nil，则执行回调（自定义回调函数）。使用sync.Map的range方法，遍历users，检测user是否还在连接中，最后心跳时间超过
设定的5min，则清除该用户。统计用户数量，发布订阅消息（那这里我就懂了，应该还是给管理员的消息后门）。

- run函数接下来的部分是管理客户端连接和断联，在客户端结构体处，一旦在写和读时发生错误，这些客户端会保存到hubc的disconnect字段，一旦此管道有值，
则会对断联客户端进行处理，关闭管道，关闭goroutine，日志记录。此处把连接的客户端和断开的客户端通过两个topic主题发布给订阅了该主题的user（大概是
管理员，由管理员订阅这个topic用来监控程序的用户情况？）

## Broadcast()
遍历guests字段，调用sendMsg函数将消息广播给访客客户端（未登陆用户？）。接着遍历user，调用sendMsg广播消息（user是app的内部用户，而client是
改用户的物理连接？所以一个user可以有多个clients，而client对应一个user。方便多端登陆管理？）

## UserLogin(uid,appId string,client *Client)
那这个函数就是注册app内部用户的方法，（appId对应不同的物理登陆设备？），app登陆之后保存到users字段中，并从访客列表清除（所以所有用户在登陆之前
都是访客？）。


# User

```go 
type User struct {
	//公共基础信息
	Uid          uint             `json:"uid"`                //整型唯一ID
	...
	
	//最后心跳时间
	LastHeartbeatTime time.Time

	//用户相关数据
	Hub        *Hubc     `json:"-"`
	AppClients []*Client `json:"-"` //appId对应客户端

	SubTopics map[string]*Topic `json:"-"` //topicId订阅的主题名称及信息
	sync.RWMutex
}
```

user结构体是用来管理app内的用户身份的结构体。内部包含该用户的多个客户端，以及各种身份信息。最主要的还是AppClients、SubTopics字段，前者用来做
多端登陆管理，后者用来给不同场景下的所有user发布消息（目前就有管理员的后门topic，disconnect...）

- **主要函数**

## AddSubTopic(topic *Topic)
给user订阅指定topic

## UnsubTopic(topicId string)
给user取消订阅topic

## UnsubAllTopics()
给user取消订阅全部topic（应该在user断联时调用）

## appLogin()
首先遍历user的客户端连接，发现同一个appid时，拿出来与新登陆的做比对。如果ws conn不同（代表同一个设备断网再次连接？），反之将改客户端加入user
的Appclients字段。

- 总结：该方法是由客户端ws连接转为app内部用户的方法。

## appLogout()
定义了一个index=-1，如果index>-1也就是在user.Appclients内部存在这么一个客户端，那么就走logout步骤。

# PubSub & Topic
```go
type PubSub struct {
    Topics        *sync.Map      //Topics map[string]*Topic //主题名称和Top对应map
    TopicMsgQueue chan *TopicMsg //主题消息队列
}

type Topic struct {
    Id          string   //订阅主题ID
    PubSub      *PubSub  //关联PubSub
    SubUsers    sync.Map //SubUsers map[string]*time.Time //订阅用户uniqueId和订阅时间
    SubHandlers sync.Map //SubHandlers map[string]func(msg *TopicMsg) //内部组件间通知
}

type TopicMsg struct {
    Ori     any    //原始数据方便订阅主题的函数处理
    TopicId string //话题ID
    Msg     []byte //消息内容，方便客户端处理
}
```

PubSub是整个hubc中的消息分发核心，主要作用是用户订阅、取消订阅Topic，并将Topic的消息分发给订阅用户，还能管理Topic（初始化、删除）

- **主要函数**

## Start()
这是整个PubSub发送消息的核心，关键是监听管道，待发送消息抵达管道，管道将消息的所属topic拿到，调用SendToSubUser函数发送msg。（topicMsg的
结构也挺有意思，一份是原始数据，可以对其进行预处理，一份是发送给客户端的json数据）

## Sub & SubFunc
```go
// Sub 订阅主题
func (a *PubSub) Sub(topicId string, user *User) {
	a.initTopic(topicId).AddSubUser(user)
}

// SubFunc 以函数方式订阅
func (a *PubSub) SubFunc(topicId string, f func(msg *TopicMsg)) {
	a.initTopic(topicId).AddSubHandle(f)
}
```

这两个函数，前者是给user订阅指定topic，后者是给指定topic增加消息的预处理方法（一个是对用户，一个是对后端）

# aqi总结
aqi是一个ws消息传输框架，服务器能够定制路由做到请求分发，匹配到对应函数执行。同时也能够发布topic给所有订阅者推送消息。既保证了客户端请求需要，
还增加了服务器的主动推送能力。前后端的消息传递统一使用json格式，好处是方便解析易于开发，可读性好。