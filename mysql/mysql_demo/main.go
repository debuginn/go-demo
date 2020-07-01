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
func queryRowDemo() (u1 *user, err error) {
	// 声明查询语句
	sqlStr := "SELECT id,name,age FROM user WHERE id = ?"
	// 声明一个 user 类型的变量
	var u user
	// 执行查询并且扫描至 u
	err = db.QueryRow(sqlStr, 1).Scan(&u.id, &u.age, &u.name)
	if err != nil {
		return nil, err
	}
	u1 = &u
	return
}

// 查询多行数据
func queryMultiRowDemo() {
	sqlStr := "SELECT id,name,age FROM user WHERE id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query data failed，err:%s\n", err)
		return
	}
	// 查询完数据后需要进行关闭数据库链接
	defer rows.Close()

	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan data failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

// 增加一行数据
func insertRowDemo() {
	sqlStr := "INSERT INTO user(name, age) VALUES(?, ?)"
	result, err := db.Exec(sqlStr, "小羽", 22)
	if err != nil {
		fmt.Printf("insert data failed, err:%v\n", err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get insert lastInsertId failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, id:%d\n", id)
}

// 更新一组数据
func updateRowDemo() {
	sqlStr := "UPDATE user SET age = ? WHERE id = ?"
	result, err := db.Exec(sqlStr, 22, 1)
	if err != nil {
		fmt.Printf("update data failed, err:%v\n", err)
		return
	}
	n, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get rowsaffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除一行数据
func deleteRowDemo() {
	sqlStr := "DELETE FROM user WHERE id = ?"
	result, err := db.Exec(sqlStr, 2)
	if err != nil {
		fmt.Printf("delete data failed, err:%d\n", err)
		return
	}
	n, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get affected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

// 主函数
func main() {
	// 初始化 MySQL
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//u1, err := queryRowDemo()
	//if err != nil {
	//	fmt.Printf("err:%s", err)
	//}
	//fmt.Printf("id:%d, age:%d, name:%s\n", u1.id, u1.age, u1.name)

	//queryMultiRowDemo()

	//insertRowDemo()

	updateRowDemo()

	//deleteRowDemo
	//deleteRowDemo()
}
