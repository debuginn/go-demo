package main

import "fmt"
import "github.com/speps/go-hashids/v2"

func main() {
	hd := hashids.NewData()
	hd.Salt = "this is my salt1"
	hd.MinLength = 5
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{1510861646})
	fmt.Println(e)
	d, _ := h.DecodeWithError(e)
	fmt.Println(d)
}
