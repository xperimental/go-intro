package main

import "fmt"

// START OMIT
type person struct {
	name string
	age  int
}

type employee struct {
	person
	pay float64
}

func main() {
	fry := person{
		name: "Fry",
		age:  29,
	}
	deliveryBoy := employee{
		person: fry,
		pay:    0,
	}
	fmt.Printf("deliveryBoy.name = %s\n", deliveryBoy.name)
}

// END OMIT
