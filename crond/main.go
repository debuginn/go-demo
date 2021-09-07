package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"sync"
	"time"
)

func main() {
	c := cron.New()

	c.AddJob("@every 1s", cron.SkipIfStillRunning(cron.DefaultLogger)(cron.FuncJob(func() {
		var wg sync.WaitGroup

		wg.Add(3)
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			fmt.Printf("1 * time.Second\n")
		}()
		go func() {
			defer wg.Done()
			time.Sleep(5 * time.Second)
			fmt.Printf("5 * time.Second\n")
		}()
		go func() {
			defer wg.Done()
			time.Sleep(20 * time.Second)
			fmt.Printf("20 * time.Second\n")
		}()
		wg.Wait()
	})))
	//
	////
	//c.AddJob("@every 1s", cron.SkipIfStillRunning(cron.DefaultLogger)(cron.FuncJob(func() {
	//	//time.Sleep(time.Second * 3)
	//
	//	fmt.Printf("SkipIfStillRunning111: %v\n", time.Now())
	//})))

	c.Start()
	defer c.Stop()

	select {}
}
