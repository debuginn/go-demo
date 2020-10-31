package main

import "fmt"

func main() {
	//fmt.Println("Hello, playground")
	//
	//defer func() {
	//	//func() {
	//	//
	//	//}()
	//	if r := recover(); r != nil {
	//		fmt.Println("cccc")
	//	}
	//}()
	//// 业务代码1...
	//
	//fmt.Println("aa")
	//panic("aaaa")
	res, err := aaa()
	if err != nil && res == nil {
		fmt.Println("aa")
		return
	}

	fmt.Println("bb")
	return
}

func aaa() (err error, aa interface{}) {
	err = fmt.Errorf("11111")
	return
}
