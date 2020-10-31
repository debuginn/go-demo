package main

import (
	"fmt"
	"sync"
)

// try Lock
type Lock struct {
	c chan struct{}
}

// 创建 锁
func newLock() Lock {
	var l Lock
	// l.c 创建一个 struct 类型通道
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}

	return l
}

// 加锁
func (l Lock) Lock() bool {
	lockRes := false

	select {
	case <-l.c:
		lockRes = true
	default:
	}

	return lockRes
}

// 释放锁
func (l Lock) Unlock() {
	l.c <- struct{}{}
}

var counter int

func main() {
	var l = newLock()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			// 获取到锁执行操作，否则则加锁失败
			if !l.Lock() {
				println("lock failed")
				return
			}

			counter++
			fmt.Printf("current counter: %d\n", counter)
			l.Unlock()
		}()
	}

	wg.Wait()
}
