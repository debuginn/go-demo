package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	// 定义模版
	// 解析模版
	t := template.New("f.tmpl")
	_, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Printf("paese template faild error: %v", err)
		return
	}
	name := "Debuginn"

	// 渲染模版
	_ = t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", f1)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Http Servertart faild, err:%v", err)
		return
	}
}
