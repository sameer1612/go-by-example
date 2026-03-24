package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string, 2)

	messages <- "message #1"
	messages <- "message #2"

	go func() {
		messages <- "message #3"
	}()

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	time.Sleep(time.Millisecond)
}
