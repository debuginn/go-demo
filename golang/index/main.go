package main

import (
	"fmt"
	"index/suffixarray"
	"sort"
)

func main() {

	source := []byte("hello world, hello china")
	index := suffixarray.New(source)

	fmt.Printf("================ %+v\n", index)
	aa := index.Bytes()
	fmt.Printf("=========aa======= %+v\n", aa)

	offsets := index.Lookup([]byte("hello"), -1)

	sort.Ints(offsets)

	fmt.Printf("%v", offsets)

}
