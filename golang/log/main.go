package main

import (
	"log"
	"os"
)

func main() {
	prefix := "[mi log info]"
	logger := log.New(os.Stdout, prefix, log.LstdFlags|log.Lshortfile)
	logger.Print("Hello World")
}
