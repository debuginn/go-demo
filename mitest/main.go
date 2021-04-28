package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

var cache = map[interface{}]interface{}{}

func keepalloc() {
	for i := 0; i < 10000; i++ {
		m := make([]byte, 1<<10)
		cache[i] = m
	}
}

func keepalloc2() {
	for i := 0; i < 100000; i++ {
		go func() {
			select {}
		}()
	}
}

func main() {

	f, _ := os.Create("trace.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()
	keepalloc()
	keepalloc2()
}

func demo2() {
	tick := time.Tick(5 * time.Second)

	for {
		<-tick

		fmt.Printf("======= \n")
	}
}

func demo1() {
	go func() {
		for {
		}
	}()
	time.Sleep(time.Millisecond)
	runtime.GC()
	println("OK")
}

func demo() {
	i := 1
	fmt.Printf("1  %p---%d \n", &i, i)

	defer func() {
		i++
		fmt.Printf("1.5  %p---%d \n", &i, i)
	}()

	defer func() {
		i++
		fmt.Printf("2  %p---%d \n", &i, i)
	}()

	defer func(i int) {
		fmt.Printf("3  %p---%d \n", &i, i)
	}(i)

	i++
	fmt.Printf("4  %p---%d \n", &i, i)
	return
}
