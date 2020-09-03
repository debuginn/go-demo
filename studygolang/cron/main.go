package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"time"
)

func main() {
	i := 0
	c := cron.New(
		cron.WithLogger(cron.VerbosePrintfLogger(log.New(os.Stdout, "cron ", log.Lshortfile))))

	_, _ = c.AddFunc("@every 30s", func() {
		i++
		fmt.Printf("task %d\n", i)
	})

	//time.ParseDuration()

	c.Start()

	//cron.Parser{}

	for {
		time.Sleep(time.Second)
	}
}
