package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(strings.NewReader("hello world !"))

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
