package main

import (
	"fmt"
)

type person struct {
	name string
	age  uint8
}

func createNewPerson(name string) *person {
	p := person{name: name}
	p.age = 1
	return &p
}

func main() {
	fmt.Println(person{"Sameer", 31})
	fmt.Println(person{name: "Jyoti", age: 29})
	fmt.Println(person{name: "Bob"})
	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(createNewPerson("Robert"))

	s := person{"Sean", 24}
	fmt.Println(s)

	sp := &s
	fmt.Println(sp.age)

	phone := struct {
		company string
		model   string
	}{
		"Apple",
		"iPhone 13 Pro",
	}
	fmt.Println(phone)
}
