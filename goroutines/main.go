package main

import (
	"fmt"
	"time"
)

func countToThree(source string) {
	for i := range 3 {
		fmt.Println(source, ":", i)
	}
}

func main() {
	countToThree("direct")

	go countToThree("goroutine#1")

	go func() {
		for range 3 {
			fmt.Println("Something random!")
		}
	}()

	go countToThree("goroutine#2")

	time.Sleep(time.Millisecond * 20)
	fmt.Println("Done")
}
