package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	env := os.Getenv("GOTRACEBACK")
	fmt.Printf("GOTRACEBACK: %s\n", env)

	for i := 0; i < 3; i++ {
		go a()
	}
	go b()
	for i := 0; i < 3; i++ {
		go a()
	}

	time.Sleep(time.Second * 1)
}

func a() {
	time.Sleep(time.Millisecond * 1)
	fmt.Printf("aaaaaaa\n")
}

func b() {
	time.Sleep(time.Millisecond * 1)
	panic("b panic ......")
}
