package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/json", func(context *gin.Context) {
		// 方法1 使用map
		//data := map[string]interface{}{
		//	"id" :   1,
		//	"name" : "debuginn",
		//	"sex" :  "men",
		//}
		data := gin.H{
			"id":   1,
			"name": "debuginn",
			"sex":  "men",
		}
		context.JSON(http.StatusOK, data)
	})

	// 方法2 结构体
	type msg = struct {
		Name    string
		Age     int
		Message string
	}
	r.GET("/another_json", func(context *gin.Context) {
		data := msg{
			Name:    "topic",
			Age:     20,
			Message: "hello",
		}
		context.JSON(http.StatusOK, data)
	})
	_ = r.Run(":9090")
}
