package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 2、解析模版
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("template import faild, err: %v", err)
		return
	}
	// 3、渲染模版
	err = t.Execute(w, "topic")
	if err != nil {
		fmt.Println("template error, err: %v", err)
	}
}

func main() {
	// 1、定义路由及端口
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP Server Start faild, err: %v", err)
		return
	}
}
