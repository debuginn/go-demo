package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()
	// Gin 框架给模版添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.Static("/xxx", "./static")
	r.LoadHTMLGlob("./templates/**/*")

	r.GET("/posts/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index.html",
		})
	})

	r.GET("/users/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "<a href=\"https://www.debuginn.cn\">debuginn</a>",
		})
	})

	r.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	_ = r.Run(":8080")
}
