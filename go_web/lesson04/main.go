package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Sex  string
	Age  int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 2、解析模版
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("Template import faild, err: %v \n", err)
		return
	}
	u1 := User{
		Name: "张三",
		Sex:  "男",
		Age:  22,
	}
	// 3、渲染模版
	err = t.Execute(w, u1)
	if err != nil {
		fmt.Printf("template error, err: %v \n", err)
	}
}

func main() {
	// 1、定义路由及端口
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP Server Start faild, err: %v \n", err)
		return
	}
}
