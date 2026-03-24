package main

import "fmt"

func doPing(pings chan<- string, msg string) {
	pings <- msg
}

// Writes message from pings channel to pongs channel
func doPongs(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	doPing(pings, "Ping msg")
	doPongs(pings, pongs)

	fmt.Println(<-pongs)
}
