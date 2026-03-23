package main

import (
	"fmt"
	"iter"
)

func genFibonacci() iter.Seq[int] {
	return func(yeild func(int) bool) {
		a, b := 0, 1

		for {
			if !yeild(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	for n := range genFibonacci() {
		if n > 10 {
			break
		}
		fmt.Println(n)
	}
}
