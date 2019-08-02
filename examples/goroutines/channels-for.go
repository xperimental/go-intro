package main

import (
	"fmt"
)

func main() {
	messages := make(chan string, 2)

	messages <- "Hello"
	messages <- "World"

	// START OMIT
	for msg := range messages { // HL
		fmt.Println(msg)
	}
	// END OMIT
}
