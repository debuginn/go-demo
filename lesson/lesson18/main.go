package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var (
	DB *gorm.DB
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

/**
连接 mysql 数据库
*/
func initMySQL() (err error) {
	dsn := "root:root@(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	// 若是连接上 为nil 否者不是
	return DB.DB().Ping()
}

func main() {
	// 连接数据库
	// 连接 mysql 数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer DB.Close() // 程序退出关闭数据库
	// 模型绑定
	DB.AutoMigrate(&Todo{})

	// 初始化 gin框架
	r := gin.Default()
	// 将静态资源替换到全局
	r.Static("/static", "static")
	// 加载 HTML 静态文件
	r.LoadHTMLGlob("templates/*")
	// 路由请求
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", func(context *gin.Context) {
			// 1、接收前端数据
			var todo Todo
			_ = context.BindJSON(&todo)
			// 2、存入数据库
			if err = DB.Create(&todo).Error; err != nil {
				context.JSON(http.StatusOK, gin.H{
					"err": err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, todo)
			}
			// 3、返回响应
		})
		// 查看全部
		v1Group.GET("/todo", func(context *gin.Context) {
			var todoList []Todo
			if err = DB.Find(&todoList).Error; err != nil {
				context.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, todoList)
			}
		})
		// 查看指定
		v1Group.GET("/todo/:id", func(context *gin.Context) {

		})
		// 修改
		v1Group.PUT("/todo/:id", func(context *gin.Context) {
			id, ok := context.Params.Get("id")
			if !ok {
				context.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				return
			}
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				context.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			_ = context.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				context.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, todo)
			}

		})
		// 删除
		v1Group.DELETE("/todo/:id", func(context *gin.Context) {
			id, ok := context.Params.Get("id")
			if !ok {
				context.JSON(http.StatusOK, gin.H{"error": "无效的 id"})
				return
			}
			if err = DB.Where("id=?", id).Delete(&Todo{}).Error; err != nil {
				context.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			} else {
				context.JSON(http.StatusOK, gin.H{id: "deleted"})
			}

		})
	}

	// 启动服务
	_ = r.Run(":9090")
}
