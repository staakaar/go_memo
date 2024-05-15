//go:build ignore

package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("file.txt")
	if err != nil {
		return fmt.Errorf("cannot open file: %v", err)
	}
	defer f.Close()

	dirname := "path/to"

	err := os.Mkdir(dirname, 0755)
	if err != nil {
		return fmt.Errorf("cannot create directory: %v", err)
	}
	defer os.RemoveAll(dirname)
}
