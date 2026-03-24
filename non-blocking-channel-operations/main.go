package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("Recieved:", msg)
	default:
		fmt.Println("No messages recieved!")
	}

	msg := "Hi"
	select {
	case messages <- msg:
		fmt.Println("Sent:", msg)
	default:
		fmt.Println("No messages sent!")
	}

	select {
	case msg := <-messages:
		fmt.Println("Recived message:", msg)
	case sig := <-signals:
		fmt.Println("Recived signal:", sig)
	default:
		fmt.Println("No activity")
	}
}
