package main

import "fmt"

func main() {
	// START OMIT
	var value int = 8
	var valuePointer *int = &value // HL
	var value2 = *valuePointer
	// END OMIT

	fmt.Printf("value = %#v\n", value)
	fmt.Printf("valuePointer = %#v\n", valuePointer)
	fmt.Printf("*valuePointer = %#v\n", *valuePointer)
}
