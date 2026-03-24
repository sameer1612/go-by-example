package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c <- "Sameer"
	}()

	select {
	case res := <-c:
		fmt.Println("Done by", res)
	case <-time.After(time.Second):
		fmt.Println("Timeout")
	}
}
