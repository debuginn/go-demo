package main

import (
	"bytes"
	"fmt"
)

func main() {

	// new reader
	fmt.Printf("new reader: %v\n", bytes.NewReader([]byte("aaabbbccc, 世界")))

	// compare
	fmt.Printf("compare: %d\n", bytes.Compare([]byte("aaabbbccc, 世界"), []byte("aaabbbccc, 世界")))

	// index
	fmt.Printf("index: %d\n", bytes.Index([]byte("aaabbb"), []byte("bb")))
	// index rune
	fmt.Printf("index rune:%d\n", bytes.IndexRune([]byte("mengxiaoyu"), 'x'))
	// index any
	fmt.Printf("index any:%d\n", bytes.IndexAny([]byte("aabbcc"), "bc"))
	//bytes.IndexAny()
	//bytes.IndexByte()
	//bytes.IndexFunc()

	// map
	rule1 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Printf("map: %s\n", bytes.Map(rule1, []byte("aaaDFGGHHVBHR")))

}
