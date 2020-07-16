package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // host:port
		Password: "",               // set password
		DB:       0,                // default db
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := initClient(); err != nil {
		fmt.Printf("init redis client failed: err:%v\n", err)
		return
	} else {
		fmt.Printf("init redis client success\n")
	}
}
