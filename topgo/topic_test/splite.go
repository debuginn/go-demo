package splite

import (
	"fmt"
	"strings"
)

func Splite(s string, sep string) (result []string) {
	i := strings.Index(s, sep)

	fmt.Println(i)

	for i > -1 {
		result = append(result, s[:i])
		s = s[i+1:]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
