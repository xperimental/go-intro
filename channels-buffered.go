package main

import (
	"fmt"
)

// START OMIT
func main() {
	messages := make(chan string, 2) // HL

	messages <- "Hello"
	messages <- "World"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

// END OMIT
