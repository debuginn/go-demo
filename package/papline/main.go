package main

import (
	"fmt"
	"math/rand"
)

func main() {

	num := rand.Int()
	fmt.Println(num)
	fmt.Println(num % 10)

	//c := gen(2, 4, 5, 6, 7)
	//out := sq(c)
	//
	//fmt.Println(<-out)
	//fmt.Println(<-out)
	//<-out
	//fmt.Println(<-out)
	//fmt.Println(<-out)
}

func gen(num ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, v := range num {
			out <- v
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
