package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path := filepath.Join(os.TempDir(), "defer.txt")
	f := createFile(path)
	defer closeFile(f)
	writeFile(f)
}

func createFile(path string) *os.File {
	fmt.Println("Creating file...")
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return file
}

func closeFile(file *os.File) {
	fmt.Println("Closing file...")
	err := file.Close()
	if err != nil {
		panic(err)
	}
}

func writeFile(file *os.File) {
	fmt.Println("Writing to file...")
	_, err := fmt.Fprintln(file, "Hello, defer!")
	if err != nil {
		panic(err)
	}
}
