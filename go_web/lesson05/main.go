package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	say := func(name string) (string, error) {
		return name + "真帅", nil
	}
	// 定义模版
	t := template.New("f.tmpl")
	// 定义模版引擎函数
	t.Funcs(template.FuncMap{
		"say": say,
	})
	// 解析模版
	_, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Printf("paese template faild error: %v", err)
		return
	}
	name := "Debuginn"

	// 渲染模版
	_ = t.Execute(w, name)
}

func demo1(w http.ResponseWriter, r *http.Request) {
	// 定义模版
	t := template.New("t.tmpl")
	// 解析模版
	_, err := t.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Printf("parse template file error:%v /n", err)
		return
	}
	// 渲染模版
	name := "Debuginn"
	_ = t.Execute(w, name)
}
func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/demo1", demo1)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Http Servertart faild, err:%v", err)
		return
	}
}
