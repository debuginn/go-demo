package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/index", func(context *gin.Context) {
		//context.JSON(http.StatusOK, gin.H{
		//	"message" : "ok",
		//})
		context.Redirect(http.StatusMovedPermanently, "https://www.debuginn.cn")
	})

	r.GET("/a", func(context *gin.Context) {
		// 跳转至 b 方法下操作
		context.Request.URL.Path = "/b"
		// 继续操作函数
		r.HandleContext(context)
	})

	r.GET("/b", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	_ = r.Run(":9090")
}
