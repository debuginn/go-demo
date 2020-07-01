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

func prepareInsert() {
	sqlStr := "INSERT INTO user(name, age) VALUES(?, ?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare sql failed, err:%v\n", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec("张三", 18)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}

	_, err = stmt.Exec("李四", 19)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}

	fmt.Printf("insert data success\n")
}
func main() {
	err := initMySQL()
	if err != nil {
		fmt.Println("init mysql failed, err%v\n", err)
		return
	}

	prepareInsert()
}
