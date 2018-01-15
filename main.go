package main

import (
	"fmt"

	"github.com/devfeel/dotweb"

	_ "github.com/yulibaozi/yulibaozi.com/conn"
	_ "github.com/yulibaozi/yulibaozi.com/repository/models"
	"github.com/yulibaozi/yulibaozi.com/routers"
)

func main() {
	fmt.Println("========yulibaozi.com========")
	app := dotweb.New()
	app.SetDevelopmentMode()
	//静态资源配置
	// app.HttpServer.ServerFile("/static/*filepath", "/Users/yulibaozi/GoWorkSpace/src/github.com/yulibaozi/yulibaozi.com/static")
	// app.HttpServer.Renderer().SetTemplatePath("/Users/yulibaozi/GoWorkSpace/src/github.com/yulibaozi/yulibaozi.com/views")
	app.HttpServer.SetEnabledListDir(false)
	// routers.InitRoute(app.HttpServer)
	app.SetEnabledLog(true)
	app.SetLogPath("logs")
	port := 8010
	routers.InitRoute(app.HttpServer)
	err := app.StartServer(port)
	if err != nil {
		fmt.Println("启动服务失败...")
	}
	fmt.Println("服务已经启动....")
}
