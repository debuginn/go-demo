package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	countNum := 0

	// 确认辅助变量是否都执行完成
	var wg sync.WaitGroup

	// wg 添加数目要和 创建的协程数量保持一致
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				countNum++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Printf("countNum: %d\n", countNum)
}
