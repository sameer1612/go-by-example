package main

import "fmt"

func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {
	q, rem := divmod(7, 3)
	fmt.Println("quotient =", q, "rem =", rem)
}
