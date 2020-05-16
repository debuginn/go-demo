package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHandler(content *gin.Context) {
	content.JSON(http.StatusOK, gin.H{
		"message": "123456",
	})
}

func m1(content *gin.Context) {
	fmt.Println("m1 in ...")
	start := time.Now()
	content.Next()
	cost := time.Since(start)
	println(cost)
	fmt.Println("m1 end...")
}

func main() {
	r := gin.Default()

	r.Use(m1)
	r.GET("/index", indexHandler)
	r.GET("/user", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "user",
		})
	})
	r.GET("/shop", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "shop",
		})
	})

	_ = r.Run(":8080")
}
