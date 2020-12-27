package main

import (
	"flag"
	"fmt"
)

func main() {
	foo := flag.String("foo", "", "Foo")
	bar := flag.String("bar", "", "Bar")

	flag.Parse()
	fmt.Print(*foo, *bar)
}
