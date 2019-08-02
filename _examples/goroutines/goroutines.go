package main

import (
	"fmt"
	"time"
)

// START OMIT
func worker(number int) {
	// <do work>
	fmt.Printf("worker: %d\n", number)
}

func main() {
	for i := 0; i < 10; i++ {
		// do something in background
		go worker(i) // HL
	}

	// wait for workers to complete
	time.Sleep(100 * time.Millisecond)
}

// END OMIT
