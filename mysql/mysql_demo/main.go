package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct {
	id   int
	age  int
	name string
}

// 初始化 MySQL 函数
func initMySQL() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/sql_test"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	return
}

// 查询一行数据
func QueryRowDemo() (u1 user, err error) {
	// 声明查询语句
	sqlStr := "select id,name,age from user where id = ?"

	var u user
	// 执行查询并且扫描至 u
	err = db.QueryRow(sqlStr, 1).Scan(&u.id, &u.age, &u.name)
	if err != nil {
		return u, err
	}
	u1 = u
	return
}

func main() {
	// 初始化 MySQL
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	u1, err := QueryRowDemo()
	if err != nil {
		fmt.Printf("err:%s", err)
	}
	fmt.Printf("id:%d, age:%d, name:%s\n", u1.id, u1.age, u1.name)
}
