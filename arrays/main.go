package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("a =", a)

	a[4] = 100
	fmt.Println("a =", a)

	fmt.Println("len = ", len(a))

	b := [5]int{2, 1, 5, 7, 6}
	fmt.Println("b =", b)

	bb := [...]int{2, 1, 5, 7, 6} // compiler counts the length
	fmt.Println("bb =", bb)

	c := [...]int{100, 3: 300, 5: 500} // index: value
	fmt.Println("c =", c)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD =", twoD)

	twoD = [2][3]int{
		{11, 22, 33},
		{44, 55, 66},
	}
	fmt.Println("twoD =", twoD)
}
