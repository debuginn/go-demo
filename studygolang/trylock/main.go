package main

import (
	"fmt"
	"sync"
)

// try Lock
type Lock struct {
	c chan struct{}
}

func newLock() Lock {
	var l Lock
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}

	return l
}

func (l Lock) Lock() bool {
	lockRes := false

	select {
	case <-l.c:
		lockRes = true
	default:
	}

	return lockRes
}

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
