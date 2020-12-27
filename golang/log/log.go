package main

import (
	"log"
	"os"
)

func main() {
	prefix := "[THIS IS THE LOG]"
	logger := log.New(os.Stdout, prefix, log.LstdFlags|log.Lshortfile)
	logger.Print("Hello World")
}
