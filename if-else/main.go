package main

import "fmt"

func main() {
	x := 12
	if x%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	if x%2 == 0 && x%3 == 0 {
		fmt.Println("Divisible by 6")
	}

	if x = 5; x%2 == 0 {
		fmt.Println("x is even")
	} else {
		fmt.Println("x is odd")
	}
}
