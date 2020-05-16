package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	newSlice := make(map[int]*int)

	// 注意此处：for range 循环的时候会创建每个元素的副本，而不是元素的引用
	for k, v := range slice {
		value := v
		newSlice[k] = &value
	}

	for key, val := range newSlice {
		fmt.Println(key, "=>", *val)
	}
}
