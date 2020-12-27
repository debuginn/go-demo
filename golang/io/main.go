package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := io.LimitReader(r, 4)
	io.Copy(os.Stdout, lr)

}
