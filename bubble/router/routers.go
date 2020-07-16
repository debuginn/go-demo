package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 初始化 gin框架
	r := gin.Default()

	// 将静态资源替换到全局
	r.Static("/static", "static")
	// 加载 HTML 静态文件
	r.LoadHTMLGlob("templates/*")
	// 路由请求
	r.GET("/", controller.IndexHandler)

	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看全部
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}

	return r
}
