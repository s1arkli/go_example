package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/wonli/aqi"
	"github.com/wonli/aqi/ws"
)

func main() {
	//app是AppConfig结构体实例
	app := aqi.Init(
		//设置配置文件名
		aqi.ConfigFile("config-test.yaml"),
		//服务名称,以及服务端口
		aqi.HttpServer("my server", "port"),
	)

	engine := gin.Default()
	engine.GET("/ws", func(c *gin.Context) {
		ws.HttpHandler(c.Writer, c.Request)
	})

	wsr := ws.NewRouter()

	var f = func(a *ws.Context) {
		fmt.Println("middleware")
	}

	wsr.Add("hi", f, func(a *ws.Context) {
		a.Send(ws.H{
			"hi": "root",
		})
	})

	group1 := wsr.Group("group1")
	group1.Add("hi", func(a *ws.Context) {
		a.Send(ws.H{
			"hi": "group1",
		})
	})

	app.WithHttpServer(engine)
	app.Start()
}
