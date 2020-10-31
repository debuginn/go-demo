package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	count uint64
}

func (c *Counter) Incr() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.count++
}

func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.count
}

func worker(counter *Counter, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Second)
	counter.Incr()
}

func main() {
	var counter Counter
	var wg sync.WaitGroup
	goroutineNum := 10

	wg.Add(goroutineNum)

	for i := 0; i < goroutineNum; i++ {
		go worker(&counter, &wg)
	}

	wg.Wait()

	fmt.Printf("count: %d\n", counter.Count())
}
