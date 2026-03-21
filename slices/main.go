package main

import (
	"fmt"
	"reflect"
	"slices"
)

func main() {
	var s []string
	fmt.Println("s =", s, s == nil, len(s))

	sl := make([]string, 3)
	fmt.Println("sl is a", reflect.TypeOf(sl))
	fmt.Println("sl =", sl, sl == nil, "len =", len(sl), "cap =", cap(sl))

	sl[0] = "a"
	sl[1] = "b"
	sl[2] = "c"
	fmt.Println("sl =", sl)

	sl = append(sl, "d")
	sl = append(sl, "e", "f")
	fmt.Println("sl =", sl, "len =", len(sl), "cap =", cap(sl))

	c := make([]string, len(sl))
	copy(c, sl)
	fmt.Println("c =", c)

	fmt.Println("First 3 elements:", c[0:3])
	fmt.Println("First 3 elements:", c[:3])
	fmt.Println("Last 3 elements:", c[len(c)-3:])

	slc1 := []string{"x", "y", "z"}
	slc2 := []string{"x", "y", "z"}

	if slices.Equal(slc1, slc2) {
		fmt.Println("slc1 == slc2")
	}

	studentMarks := make([][]int, 3)
	for i := range 3 {
		marksLen := i + 1
		marks := make([]int, marksLen)
		for j := range marksLen {
			marks[j] = j
		}
		studentMarks[i] = marks
	}
	fmt.Println("studentMarks =", studentMarks)
}
