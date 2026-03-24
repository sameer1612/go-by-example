package main

import (
	"fmt"
	"time"
)

func doAsyncWork(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second)
	fmt.Println("Done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go doAsyncWork(done)

	<-done
}
