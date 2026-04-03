package main

import "fmt"

func mayPanic() {
	panic("something went wrong")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	mayPanic()
	fmt.Println("After the panic...")
}
