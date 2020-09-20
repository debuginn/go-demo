package main

import (
	"fmt"
	"strconv"
)

func stringSlice2Int64Slice(strSlice []string) (int64Slice []int64) {
	for _, v := range strSlice {
		int64Res, _ := strconv.ParseInt(v, 10, 64)
		int64Slice = append(int64Slice, int64Res)
	}
	return
}

func main() {
	str := "TW1VeVdXMWlOalZhYnpaQmFUa3ZWRGt2U0ZKaFp6MDk="
	str1 := str
	fmt.Println(str1)
}
