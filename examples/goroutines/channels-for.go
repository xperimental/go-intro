package main

import (
	"fmt"
	"math/rand"
	"time"
)

var randomValue = rand.New(rand.NewSource(time.Now().Unix())).Intn

// generator writes a number of random numbers to a channel.
func generator(ch chan int, count int) {
	for i := 0; i < count; i++ {
		ch <- randomValue(100)
	}
	close(ch)
}

// START OMIT
// createGenerator returns a read-only channel with some random integers.
func createGenerator(count int) <-chan int { // HL
	ch := make(chan int)
	// start generating numbers
	go generator(ch, count)
	return ch
}

func main() {
	numbersChannel := createGenerator(10)

	for n := range numbersChannel { // HL
		fmt.Printf("We got a %d!\n", n)
	}
}

// END OMIT
