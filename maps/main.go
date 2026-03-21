package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)
	m["a"] = 5
	m["b"] = 10
	fmt.Println("m =", m, "len(m) =", len(m))
	fmt.Println("m[a] =", m["a"])

	delete(m, "b")
	fmt.Println("m =", m, "len(m) =", len(m))

	clear(m)
	fmt.Println("m =", m, "len(m) =", len(m))

	_, found := m["a"]
	if !found {
		fmt.Println("Missing key:", "a")
	}

	n1 := map[string]int{"x": 1, "y": 2}
	n2 := map[string]int{"y": 2, "x": 1}
	fmt.Println("n1 == n2 is", maps.Equal(n1, n2))
}
