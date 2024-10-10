package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bytepunk/gin-template/global"
	"github.com/bytepunk/gin-template/router"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	engine:=gin.New()
	engine.Use(gin.Recovery())
	server:=&http.Server{
		Addr: global.GlobalConfig.App.Addr,
		Handler: engine,

	}

	registerRouters(engine)
	printServerInfo()
	err:=server.ListenAndServe()
	if err!=nil {
		panic("start failed:"+err.Error())
	}
}

func registerRouters(engine *gin.Engine) {
	router.RegisterRouter(engine)
}

func printServerInfo() {
	appConfig := global.GlobalConfig.App
	fmt.Printf("\n【 当前环境: %s 当前版本: %s 接口地址: http://%s 启动时间:%v 】\n",
		appConfig.Env, appConfig.Version, appConfig.Addr,time.Now().Format("2006-01-02 15:04:05"))
}