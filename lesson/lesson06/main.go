package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 首页方法定义
func index(w http.ResponseWriter, r *http.Request) {
	// 定义模版
	t := template.New("index.tmpl")
	// 解析模版
	_, err := t.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("Parse template file error:%v\n", err)
		return
	}
	// 渲染模版
	msg := "这是首页模版"
	_ = t.Execute(w, msg)
}

//Home 方法定义
func home(w http.ResponseWriter, r *http.Request) {
	// 定义模版
	t := template.New("home.tmpl")
	// 解析模版
	_, err := t.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Printf("Parse template file error:%v\n", err)
		return
	}
	// 渲染模版
	msg := "这是 home 模版"
	_ = t.Execute(w, msg)
}

func index2(w http.ResponseWriter, r *http.Request) {
	// 定义模版
	t, err := template.ParseFiles("./templates/bash.tmpl", "./templates/index2.tmpl")
	// 解析模版
	if err != nil {
		fmt.Printf("Parse template file error:%v\n", err)
		return
	}
	// 渲染模版
	name := "index2"
	_ = t.Execute(w, name)
}
func home2(w http.ResponseWriter, r *http.Request) {
	// 定义模版
	t, err := template.ParseFiles("./templates/bash.tmpl", "./templates/home2.tmpl")
	// 解析模版
	if err != nil {
		fmt.Printf("Parse template file error:%v\n", err)
		return
	}
	// 渲染模版
	name := "home2"
	_ = t.Execute(w, name)
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Http Server start error:%v\n", err)
		return
	}
}
