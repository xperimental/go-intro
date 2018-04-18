package main

import "fmt"

func main() {
	// START OMIT
	age := map[string]int{
		"Fry":        1044,
		"Farnsworth": 177,
	}
	// END OMIT
	fmt.Printf("age = %#v\n", age)
}
