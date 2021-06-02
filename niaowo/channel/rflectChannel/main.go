package main

import (
	"fmt"
	"reflect"
)

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	var cases = createCases(ch1, ch2)
	// 执行10次
	for i := 0; i < 10; i++ {
		chosen, recv, ok := reflect.Select(cases)
		if recv.IsValid() { // recv case
			fmt.Println("recv:", cases[chosen].Dir, recv, ok)
		} else { // send case
			fmt.Println("send:", cases[chosen].Dir, ok)
		}
	}
}

// createCases _
func createCases(chs ...chan int) []reflect.SelectCase {
	var cases []reflect.SelectCase

	// create recv case
	for _, ch := range chs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}

	// create send case
	for i, ch := range chs {
		v := reflect.ValueOf(i)
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(ch),
			Send: v,
		})
	}
	return cases
}
