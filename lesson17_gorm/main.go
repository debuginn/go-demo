package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	ID    uint
	Name  string
	Age   int
	Hobby string
}

func main() {
	// 连接 mysql 数据库
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 自动迁移
	db.AutoMigrate(&UserInfo{})

	// 插入
	//u1 := UserInfo{1, "七米", 22, "篮球"}
	//db.Create(&u1)

	// 查询
	var user UserInfo
	db.First(&user)
	fmt.Printf("%v\n", user)

	// 更新
	db.Model(&user).Update("hobby", "唱 跳 rap")

	// 删除
	db.Delete(&user)
}
