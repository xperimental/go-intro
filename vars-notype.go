package main

import (
	"fmt"
)

// START OMIT
var str = "string" // will be string
var integer = 5    //         int
var float = 3.4    //         float64
// END OMIT

func main() {
	fmt.Printf("str = %t\n", str)
	fmt.Printf("integer = %t\n", integer)
	fmt.Printf("float = %t\n", float)
}
