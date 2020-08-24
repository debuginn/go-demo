package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandleGet(t *testing.T) {
	api := "http://localhost:8081/"
	res, err := http.Get(api)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		fmt.Printf("get api success")
	}
}

func GetApi() {
	api := "http://localhost:8081/"
	res, err := http.Get(api)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		fmt.Printf("get api success\n")
	}
}

func Benchmark_Main(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetApi()
	}
}
