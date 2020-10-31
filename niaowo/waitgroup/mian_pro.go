package main

import "fmt"

func main() {
	var uid int64
	uid = 11111
	isNext := false
	uids := []int64{11111, 22222, 33333}

	if InIntSlice(uids, uid) {
		isNext = true
	}

	if uid > 33333 && uid < 33337 {
		isNext = true
	}

	if !isNext {
		fmt.Printf("===================")
		return
	}

	fmt.Printf("~~~~~~~~~~~~~~~~~~~~")

}

// InIntSlice item是否存在于slice中
func InIntSlice(slice []int64, item int64) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
