package main

import "fmt"

type vehicle struct {
	brand string
	model string
}

func (v vehicle) name() string {
	return fmt.Sprintf("%s %s", v.brand, v.model)
}

type car struct {
	vehicle
	wheelCount uint8
}

type plane struct {
	vehicle
	wings uint8
}

func main() {
	c := car{
		vehicle: vehicle{
			brand: "Lamborgini",
			model: "Galardo",
		},
		wheelCount: 4,
	}

	p := plane{
		vehicle{
			"Airbus",
			"A320",
		},
		2,
	}

	fmt.Println(c.name(), "has", c.wheelCount, "wheels.")
	fmt.Println(p.name(), "has", p.wings, "wings.")
}
