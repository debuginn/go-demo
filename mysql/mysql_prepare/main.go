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

// 预处理插入数据
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

// 预处理更新数据
func prepareUpdate() {
	sqlStr := "UPDATE user SET age = ? WHERE id = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare sql failed, err:%v\n", err)
		return
	}
	_, err = stmt.Exec(18, 2)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	fmt.Printf("prepare update data success")
}

// 预处理查询数据
func prepareQuery() {
	sqlStr := "SELECT id,name,age FROM user WHERE id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare sql failed, err:%v\n", err)
		return
	}
	rows, err := stmt.Query(1)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan data failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d, name:%s, age:%d\n", u.id, u.name, u.age)
	}
}

// 预处理删除数据
func prepareDelete() {
	sqlStr := "DELETE FROM user WHERE id = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare sql failed, err:%v\n", err)
		return
	}
	result, err := stmt.Exec(3)
	n, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("delete rows failed, err:%v\n", err)
		return
	}
	if n > 0 {
		fmt.Printf("delete data success")
	} else {
		fmt.Printf("delete data error")
	}
}

func main() {
	err := initMySQL()
	if err != nil {
		fmt.Printf("init mysql failed, err%v\n", err)
		return
	}

	//prepareInsert()

	//prepareUpdate()

	prepareQuery()

	//repareDelete()
}
