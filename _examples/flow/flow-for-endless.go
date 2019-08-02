package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	for {
		fmt.Println("Loop!")
		time.Sleep(time.Second)
	}
	// END OMIT
}
