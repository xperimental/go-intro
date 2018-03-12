package main

import "fmt"

func main() {
	// START OMIT
	intArray := [5]int{1, 2, 3, 4, 5}
	firstSlice := intArray[:2]  // will be {1,2}
	secondSlice := intArray[2:] // will be {3,4,5}
	firstSlice = append(firstSlice, 6)
	// END OMIT

	fmt.Printf("intArray = %#v\n", intArray)
	fmt.Printf("firstSlice = %#v\n", firstSlice)
	fmt.Printf("secondSlice = %#v\n", secondSlice)
}
