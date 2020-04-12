package controller

import (
	"bubble/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 首页 HTML 模版渲染
func IndexHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", nil)
}

// 创建一个 Todo
func CreateTodo(context *gin.Context) {
	// 1、接收前端数据
	var todo model.Todo
	_ = context.BindJSON(&todo)
	// 2、存入数据库
	err := model.CreateATodo(&todo)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, todo)
	}
	// 3、返回响应
}

// 获取 Todo 列表
func GetTodoList(context *gin.Context) {
	todoList, err := model.GetAllTodo()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, todoList)
	}
}

// 更新一个 Todo
func UpdateATodo(context *gin.Context) {
	id, ok := context.Params.Get("id")
	if !ok {
		context.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	todo, err := model.GetATodo(id)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	_ = context.BindJSON(&todo)
	err = model.SaveATodo(todo)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, todo)
	}
}

// 删除一个 Todo
func DeleteATodo(context *gin.Context) {
	id, ok := context.Params.Get("id")
	if !ok {
		context.JSON(http.StatusOK, gin.H{"error": "无效的 id"})
		return
	}
	if err := model.DeleteATodo(id); err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
