package main

import (
	"fmt"
	"path"
)

func main() {
	p := "foo/bar.tar"

	sDir, sBase := path.Split(p)

	fmt.Println(sDir)
	fmt.Println(sBase)
	fmt.Println(path.Ext(p))
}
