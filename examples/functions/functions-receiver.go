package main

import "fmt"

// START OMIT
type person struct {
	name string
}

func (p person) Say(text string) { // p is the receiver of type person // HL
	fmt.Printf("%s says %q", p.name, text)
}

func main() {
	fry := person{
		name: "Fry",
	}

	fry.Say("Hello!")
}

// END OMIT
