package main

import (
	"io"
	"net/http"

	"github.com/wonli/aqi"
)

func main() {
	//app是AppConfig结构体实例
	app := aqi.Init(
		//设置配置文件名
		aqi.ConfigFile("config-test.yaml"),
		//服务名称,以及服务端口
		aqi.HttpServer("my server", "port"),
	)

	//路由器
	mux := http.NewServeMux()
	mux.HandleFunc("/test", Handle)

	app.WithHttpServer(mux)
	app.Start()
}

func Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("get server running"))
	case "POST":
		bt, _ := io.ReadAll(r.Body)
		w.Write(bt)
	}
}
