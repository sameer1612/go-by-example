package main

import "fmt"

// recursive function
func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	fmt.Println("Factorial of 5 =", factorial(5))

	// annonymous recursive function
	var fibonacci func(n int) int
	fibonacci = func(n int) int {
		if n < 2 {
			return n
		}
		return fibonacci(n-1) + fibonacci(n-2)
	}

	fmt.Println("7th fibonacci =", fibonacci(7))
}
