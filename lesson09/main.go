package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/web", func(context *gin.Context) {
		// 获取浏览器发出请求携带的参数 query string
		data := context.Query("query")
		context.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	_ = r.Run(":9090")
}
