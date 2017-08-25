package main

import "fmt"

// START OMIT
type person struct {
	name string
}

func (p person) SayHello() {
	fmt.Printf("%s says hello!", p.name)
}

// END OMIT

func main() {
	fry := person{
		name: "Fry",
	}

	fry.SayHello()
}
