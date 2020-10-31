package main

import (
	"fmt"
	"sync"
)

// 线程安全的计数器
type Counter struct {
	CounterType int
	Name        string

	mu    sync.Mutex
	count uint64
}

// 加一方法
func (c *Counter) Incr() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// 取数值方法 线程也需要受保护
func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	// 定义一个计数器
	var counter Counter

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Incr()
			}
		}()
	}
	wg.Wait()

	fmt.Printf("%d", counter.Count())
}
