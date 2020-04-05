package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"msg": "/user/index get"})
		})
		userGroup.POST("/index", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"msg": "/user/index post"})
		})
	}

	bookGroup := r.Group("/book")
	{
		bookGroup.GET("/index", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"msg": "/book/index get"})
		})
	}

	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{"404": "not found"})
	})

	_ = r.Run(":9090")
}
