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

// 事务更新操作
func transActionUpdate() {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			_ = tx.Rollback()
		}
		fmt.Printf("begin trans action failed, err:%v\n", err)
		return
	}
	sqlStr1 := "UPDATE user SET age = ? WHERE id = ?"
	result1, err := tx.Exec(sqlStr1, 20, 1)
	if err != nil {
		_ = tx.Rollback()
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	n1, err := result1.RowsAffected()
	if err != nil {
		_ = tx.Rollback()
		fmt.Printf("exec result1.RowsAffected() failed, err:%v\n", err)
		return
	}
	sqlStr2 := "UPDATE user SET age = ? WHERE id = ?"
	result2, err := tx.Exec(sqlStr2, 20, 6)
	if err != nil {
		_ = tx.Rollback()
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	n2, err := result2.RowsAffected()
	if err != nil {
		_ = tx.Rollback()
		fmt.Printf("exec result1.RowsAffected() failed, err:%v\n", err)
		return
	}

	if n1 == 1 && n2 == 1 {
		_ = tx.Commit()
		fmt.Printf("transaction commit success\n")
	} else {
		_ = tx.Rollback()
		fmt.Printf("transaction commit error, rollback\n")
		return
	}
}

func main() {
	err := initMySQL()
	if err != nil {
		fmt.Printf("init mysql failed, err%v\n", err)
		return
	}

	transActionUpdate()
}
