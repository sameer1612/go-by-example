package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	fmt.Println("Sum =", sum)

	fruits := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range fruits {
		fmt.Printf("key = %s and value = %s\n", k, v)
	}

	for i, ch := range "golang" {
		fmt.Println(i, ch)
	}
}
