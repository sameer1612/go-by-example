package main

import "fmt"

type ServerState uint8

const (
	Idle ServerState = iota
	Connected
	Stopped
)

var status = map[ServerState]string{
	Idle:      "idle",
	Connected: "connected",
	Stopped:   "stopped",
}

func main() {
	var currentState ServerState = Connected
	fmt.Println("state:", status[currentState])

	currentState = Idle
	fmt.Println("state:", status[currentState])

	currentState = Stopped
	fmt.Println("state:", status[currentState])
}
