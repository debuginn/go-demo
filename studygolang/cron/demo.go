package main

import (
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

type delayJob struct {
	count int
}

func (d *delayJob) Run() {
	time.Sleep(2 * time.Second)
	d.count++
	log.Printf("%d: hello world\n", d.count)
}

func main() {
	c := cron.New()
	_, _ = c.AddJob("@every 1s", cron.NewChain(cron.SkipIfStillRunning(cron.DefaultLogger)).Then(&delayJob{}))
	c.Start()

	time.Sleep(10 * time.Second)
}
