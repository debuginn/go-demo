package main

import "fmt"

func main() {
	// 针对于 new 创建一个 list，不能使用 append 追加
	// 只能对于 创建 list 进行 make 操作
	// ./main.go:7:15: first argument to append must be slice; have *[]int
	list := new([]int)
	list = append(list, 1)
	fmt.Println(list)
}
