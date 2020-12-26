package main

import (
	"fmt"
)

func main() {
	n := []float32{120.4, -46.7, 32.50, 34.65, -67.45}
	fmt.Printf("The total is %.02f\n", sum(n))
}

func sum(s []float32) float32 {
	var t float32
	for _, v := range s {
		if t < 0 {
			t = Add(t, v)
		} else {
			t = Sub(t, v)
		}
	}

	return t
}

func Add(a, b float32) float32 {
	return a + b
}

func Sub(a, b float32) float32 {
	return a - b
}
