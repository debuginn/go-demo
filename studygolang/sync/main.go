package main

import (
	"fmt"
	"sync"
	"time"
	//"time"
)

var counter int

func main() {
	beginTime := time.Now()
	//fmt.Printf("----------- sync wait group -------------\n")
	//syncDemo()
	fmt.Printf("----------- sync wait group lock -------------\n")
	syncDemoLock()

	endTime := time.Since(beginTime)
	fmt.Printf("time spend: %s\n", endTime)
}

func syncDemo() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			counter++
		}()
	}

	wg.Wait()

	println(counter)
}

func syncDemoLock() {
	var wg sync.WaitGroup
	var l sync.Mutex

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			l.Lock()
			counter++
			l.Unlock()
		}()
	}

	wg.Wait()

	println(counter)
}
