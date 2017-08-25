package main

import "fmt"

// START OMIT
type person struct {
	name string
	age  int
}

func main() {
	fry := person{
		name: "Fry",
		age:  29,
	}
	fmt.Printf("fry = %#v", fry)
}

// END OMIT
