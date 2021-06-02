package topic05

import "fmt"

type AA struct {
	aa string
	bb bool
}

func get() map[string]*AA {
	m := make(map[string]*AA)
	s := []AA{
		{"aaa", true},
		{"bbb", true},
	}

	for _, v := range s {
		m[v.aa] = &v
	}
	return m
}

func Cat() {
	b := get()
	for k, b := range b {
		fmt.Printf("key:%v==v:%v\n", k, b)
	}
}
