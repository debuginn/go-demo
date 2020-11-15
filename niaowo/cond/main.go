package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	var ready int

	for i := 0; i < 10; i++ {
		go func(i int) {
			// 随机等待 10s
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

			c.L.Lock()
			ready++
			c.L.Unlock()

			log.Printf("goroutine #%d ready\n", i)

			// 广播
			c.Broadcast()
		}(i)
	}

	c.L.Lock()
	for ready != 10 {
		c.Wait()
		log.Println("唤醒一次")
	}
	c.L.Unlock()

	// 所有准备就绪
	log.Println("all goroutine ready... go")
}
