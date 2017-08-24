package main

import "fmt"

// START OMIT
func Public() {
	fmt.Println("Can be called from outside.")
}

func private() {
	fmt.Println("Can only be called from inside package.")
}

var (
	PublicValue  = 42
	privateValue = 5
)

// END OMIT
