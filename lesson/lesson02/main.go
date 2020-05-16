package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(content *gin.Context) {
	content.JSON(200, gin.H{
		"message": "Hello world",
	})
}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 指定用户使用 Get进行操作
	r.GET("/hello", sayHello)

	r.GET("/book", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"method": "get",
		})
	})
	r.POST("/book", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"method": "post",
		})
	})
	r.PUT("/book", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"method": "put",
		})
	})
	r.DELETE("/book", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"method": "delete",
		})
	})

	// 执行服务
	_ = r.Run()
}
