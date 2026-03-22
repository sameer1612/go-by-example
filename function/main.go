package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func multiply(a, b, c int) int {
	return a * b * c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1 + 2 =", res)

	res = multiply(1, 2, 3)
	fmt.Println("1 * 2 * 3 =", multiply(1, 2, 3))
}
