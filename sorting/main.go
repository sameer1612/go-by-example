package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	ints := []int{3, 1, 2}
	slices.Sort(ints)
	fmt.Println("Ints:", ints)

	fmt.Println("Ints sorted:", slices.IsSorted(ints))
}
