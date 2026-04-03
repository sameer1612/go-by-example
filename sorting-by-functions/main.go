package main

import (
	"cmp"
	"fmt"
	"slices"
)

func lenCmp(a, b string) int {
	return cmp.Compare(len(a), len(b))
}

type Person struct {
	name string
	age  int
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	people := []Person{
		{"Jax", 37},
		{"Maya", 28},
		{"Alex", 30},
	}

	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(people)
}
