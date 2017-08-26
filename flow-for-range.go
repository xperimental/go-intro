package main

import "fmt"

func main() {
	// START OMIT
	names := []string{"Fry", "Leela", "Farnsworth"}
	for i, v := range names {
		fmt.Printf("index: %d value: %s\n", i, v)
	}
	// END OMIT
}
