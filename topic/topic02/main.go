package main

import "fmt"

func main() {
	// 使用 append 向 slice 添加元素，是追加在后面，并非是从0开始
	s := make([]int, 2)
	s = append(s, 1, 2, 3, 4, 5, 6)
	fmt.Println(s)
}
