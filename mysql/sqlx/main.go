package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type user struct {
	Id   int    `db:"id"`
	Age  int    `db:"age"`
	Name string `db:"name"`
}

// 初始化数据库
func initMySQL() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/sql_test"
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("connect server failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)
	return
}

// 查询一行数据
func queryRow() {
	sqlStr := "SELECT id, name, age FROM user WHERE id = ?"

	var u user
	if err := db.Get(&u, sqlStr, 1); err != nil {
		fmt.Printf("get data failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d, name:%s, age:%d\n", u.Id, u.Name, u.Age)
}

// 查询多行数据
func queryMultiRow() {
	sqlStr := "SELECT id, name, age FROM user WHERE id > ?"
	var users []user
	if err := db.Select(&users, sqlStr, 0); err != nil {
		fmt.Printf("get data failed, err:%v\n", err)
		return
	}

	for i := 0; i < len(users); i++ {
		fmt.Printf("id:%d, name:%s, age:%d\n", users[i].Id, users[i].Name, users[i].Age)
	}
}

// 插入数据
func insertRow() {
	sqlStr := "INSERT INTO user(name, age) VALUES(?, ?)"
	result, err := db.Exec(sqlStr, "Meng小羽", 22)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get insert id failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert data success, id:%d\n", insertID)
}

// 更新数据
func updateRow() {
	sqlStr := "UPDATE user SET age = ? WHERE id = ?"
	result, err := db.Exec(sqlStr, 22, 6)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get affected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update data success, affected rows:%d\n", affectedRows)
}

// 删除一行
func deleteRow() {
	sqlStr := "DELETE FROM user WHERE id = ?"
	result, err := db.Exec(sqlStr, 4)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get affected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete data success, affected rows:%d\n", affectedRows)
}

func main() {
	if err := initMySQL(); err != nil {
		panic(err)
	}
	//queryRow()
	//insertRow()
	//updateRow()
	deleteRow()
	queryMultiRow()
}
