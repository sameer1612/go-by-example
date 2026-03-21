package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	switch i {
	case 1:
		fmt.Println("true")
	case 0:
		fmt.Println("false")
	default:
		fmt.Println("idk")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	whatAmI := func(i any) {
		switch t := i.(type) {
		case bool:
			fmt.Println("It's a bool")
		case int:
			fmt.Println("It's an int")
		default:
			fmt.Printf("IDK what is a %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(10)
	whatAmI("sameer")
	whatAmI(func() {})
}
