package main

import "fmt"

// START OMIT
var (
	IsVisibleFromOutside       = true
	canOnlyBeUsedInsidePackage = 42
)

func Public() {
	fmt.Println("Can be called from outside.")
}

func private() {
	fmt.Println("Can only be called from inside package.")
}

// END OMIT
