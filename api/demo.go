package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld(ctx *gin.Context) {
	data:=make(map[string]string)
	data["data"]="hello world"
	ctx.JSON(http.StatusOK, data)
}