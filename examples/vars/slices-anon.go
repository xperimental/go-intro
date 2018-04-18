package main

import "fmt"

func main() {
	// START OMIT
	ints := []int{0, 1, 2} // directly work with slice
	ints = append(ints, 3)
	// END OMIT

	fmt.Printf("ints = %#v\n", ints)
}
