package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/user/:name/:age", func(context *gin.Context) {
		// 获取路径参数
		name := context.Param("name")
		age := context.Param("age")
		context.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.GET("/datetime/:year/:month", func(context *gin.Context) {
		year := context.Param("year")
		month := context.Param("month")

		context.JSON(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
		})
	})

	_ = r.Run(":9090")
}
