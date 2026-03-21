package main

import "fmt"

func main() {
	i := 1

	// like while
	for i <= 3 {
		fmt.Println("i =", i)
		i += 1
	}

	// classic for
	for j := 0; j < 5; j++ {
		fmt.Println("j =", j)
	}

	// python like range
	for k := range 3 {
		fmt.Println("k =", k)
	}

	// continue
	for l := range 10 {
		if l%2 == 0 {
			continue
		}
		fmt.Println("l =", l)
	}
}
