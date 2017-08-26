package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	// Create channel for passing int values
	ch := make(chan int) // HL

	go func() {
		time.Sleep(1 * time.Second)
		ch <- 21 + 21
	}()

	fmt.Println("Waiting for result...")
	value := <-ch
	fmt.Printf("Result: %d\n", value)
}

// END OMIT
