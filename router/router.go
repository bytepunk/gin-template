package router

import (
	"github.com/bytepunk/gin-template/api"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(engine *gin.Engine) {
	demoRouter:=engine.Group("demo")
	demoRouter.GET("hello", api.HelloWorld)
}