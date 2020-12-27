package main

import (
	"path/filepath"
	"os"
	"fmt"
)

func main() {

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(info.Name())
		return nil
	})
}
