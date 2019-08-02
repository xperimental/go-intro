package main

import "fmt"

func main() {
	// START OMIT
	var fiveIntArray = [5]int{1, 2, 3, 4, 5}
	var intSlice = fiveIntArray[:2] // first two elements
	//intSlice[1] = 5               // modifies original element
	// END OMIT

	fmt.Printf("fiveIntArray = %#v\n", fiveIntArray)
	fmt.Printf("intSlice = %#v\n", intSlice)
	fmt.Printf("len = %d\n", len(intSlice))
	fmt.Printf("cap = %d\n", cap(intSlice))
}
