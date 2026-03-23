package main

import "fmt"

func indexOf[S ~[]E, E comparable](s S, e E) int {
	for i := range s {
		if s[i] == e {
			return i
		}
	}

	return -1
}

type valueObject[T interface{ ~string | ~int }] struct {
	value T
}

func main() {
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println("Index of banana:", indexOf(fruits, "banana"))

	s := valueObject[string]{"sameer"}
	fmt.Println("s =", s.value)

	i := valueObject[int]{5}
	fmt.Println("i =", i.value)
}
