package main

import (
	"os"
	"path/filepath"
)

func main() {
	// panic("Something went wrong!")

	path := filepath.Join(os.TempDir(), "file.txt")
	_, err := os.Open(path)
	if err != nil {
		panic(err)
	}
}
