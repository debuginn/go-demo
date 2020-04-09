package main

import (
	"bubble/dao"
	"bubble/model"
	"bubble/router"
)

func main() {
	// 连接数据库
	// 连接 mysql 数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.CloseMySQL() // 程序退出关闭数据库
	// 模型绑定
	dao.DB.AutoMigrate(&model.Todo{})
	// 注册路由
	r := routers.SetupRouter()
	// 启动服务
	_ = r.Run(":9090")
}
