package generics

import (
	"fmt"
	"testing"
)

func TestMapKeys(t *testing.T) {

	mapA := make(map[string]string, 0)
	mapA["a"] = "a"
	mapA["b"] = "b"
	mapA["c"] = "c"

	a := MapKeys(mapA)
	fmt.Printf("=======%+v\n", a)
}
