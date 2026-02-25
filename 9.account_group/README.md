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
