package main

import "fmt"

func main() {
	// START OMIT
	var valuePointer *int = nil
	// END OMIT

	fmt.Printf("valuePointer = %#v\n", valuePointer)
}
