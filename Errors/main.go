package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("Divide by 0 is not allowed!")

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, DivideByZeroError
	}

	return a / b, nil
}

func main() {
	q, err := divide(5, 0)
	if err == nil {
		fmt.Println("q =", q)
	} else {
		if errors.Is(err, DivideByZeroError) {
			fmt.Println("You got a divide by zero error!")
		}
	}
}
