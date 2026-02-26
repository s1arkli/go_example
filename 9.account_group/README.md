# Group
```go
    account := v1.Group("/account")
```
作为account路由分组，其下包括登陆、注销、恢复、详情接口，满足登陆项目的各种需求。

## login
```go
    account.POST("/login", ui.Login)
```
- 注册路由的入口，当访问到../login路由时，触发ui.Login函数。

```go
type LoginReq struct {
	//手机号，手机号登陆必传
	Phone string `json:"phone"`
	//验证码，手机号登陆必传
	Code string `json:"code"`
	//易盾token，一键登陆必传
	Token string `json:"token"`
	//运营商token，一键登陆必传
	AccessToken string `json:"access_token"`
}
```
- 登陆请求结构体中定义四个字段，其中分为两种登陆方式，手机号验证码形式登陆以及易盾第三方一键登陆。

## Service Login

---
- 一键登录
    - 需要两个token，对易盾发起post请求，拿到用户手机号。
    - 检查手机号是否为已注册用户，如果不是，自动注册。
    - 生成限时7天token，内部保存userID信息（用于绑定该用户，使用此token就默认是该userID的用户），token返回给客户端
---
- 手机号+验证码登录
    - 在登录之前，需要向阿里云发送post请求，给该用户手机号发送验证码短信（验证码由go服务生成，阿里云只负责发送短信）。同时将手机号作为key储存
      验证码。
    - 收到手机号+验证码，读取缓存验证手机号和其验证码是否对应上。

## Auth middleware

```go
func Auth(ac *jwt.AuthClaims, account *repository.Account) gin.HandlerFunc {
	return func(c *gin.Context){}
}
```
**账号认证的真正实现**，这是一个token认证中间件，主要原理是校验header中的token是否valid，然后将token中的userID作为身份认证做后续传递。

- 从header中拿到token，解析token，拿到其中保存的userID作为该用户的身份信息往后传递。同时也会对userID进行检验，看看是否是数据库中存在的
合法用户。
- 把解析的userID放到gin.Context中，统一使用getUserID函数获取。
```go
func getUserID(c *gin.Context) int64 {
	return c.GetInt64("user_id")
}
```