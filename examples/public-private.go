package main

import "fmt"

// START OMIT
func Public() {
	fmt.Println("Can be called from outside.")
}

func private() {
	fmt.Println("Can only be called from inside package.")
}

// END OMIT
