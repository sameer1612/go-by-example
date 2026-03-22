package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for num := range nums {
		total += num
	}

	return total
}

func main() {
	fmt.Println("sum =", sum(1, 2, 3, 4, 5))
	fmt.Println("sum =", sum([]int{1, 2, 3}...))
}
