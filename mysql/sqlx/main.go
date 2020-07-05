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

// 使用 named 方法插入数据
func insertNamedExec() {
	sqlStr := "INSERT INTO user(name, age) VALUES(:name, :age)"
	result, err := db.NamedExec(sqlStr, map[string]interface{}{
		"name": "里斯",
		"age":  18,
	})
	if err != nil {
		fmt.Printf("named exec failed, err:%v\n", err)
		return
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get last insert id failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert data success, id:%d\n", insertId)
}

// 绑定查询
func selectNamedQuery() {
	sqlStr := "SELECT id, name, age FROM user WHERE age = :age"
	rows, err := db.NamedQuery(sqlStr, map[string]interface{}{
		"age": 22,
	})
	if err != nil {
		fmt.Printf("named query failed failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		if err := rows.StructScan(&u); err != nil {
			fmt.Printf("struct sacn failed, err:%v\n", err)
			continue
		}
		fmt.Printf("%#v\n", u)
	}
}

// 事务操作
func updateTransaction() (err error) {
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("transaction begin failed, err:%v\n", err)
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			fmt.Printf("transaction rollback")
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
			fmt.Printf("transaction commit")
			return
		}
	}()

	sqlStr1 := "UPDATE user SET age = ? WHERE id = ? "
	reuslt1, err := tx.Exec(sqlStr1, 18, 1)
	if err != nil {
		fmt.Printf("sql exec failed, err:%v\n", err)
		return err
	}
	rows1, err := reuslt1.RowsAffected()
	if err != nil {
		fmt.Printf("affected rows is 0")
		return
	}
	sqlStr2 := "UPDATE user SET age = ? WHERE id = ? "
	reuslt2, err := tx.Exec(sqlStr2, 19, 5)
	if err != nil {
		fmt.Printf("sql exec failed, err:%v\n", err)
		return err
	}
	rows2, err := reuslt2.RowsAffected()
	if err != nil {
		fmt.Printf("affected rows is 0\n")
		return
	}

	if rows1 > 0 && rows2 > 0 {
		fmt.Printf("update data success\n")
	}
	return
}

func main() {
	if err := initMySQL(); err != nil {
		panic(err)
	}
	//queryRow()
	//insertRow()
	//updateRow()
	//deleteRow()
	//insertNamedExec()
	//selectNamedQuery()
	queryMultiRow()
	//_ = updateTransaction()
}
