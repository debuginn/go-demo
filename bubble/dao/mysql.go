package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

// 连接 mysql 数据库
func InitMySQL() (err error) {
	dsn := "root:root@(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	// 若是连接上 为nil 否者不是
	return DB.DB().Ping()
}

// 关闭数据库
func CloseMySQL() {
	DB.Close()
}
