package main

import "fmt"

func main() {
	// START OMIT
	intSlice := []int{}
	for i := 0; i < 10; i++ {
		intSlice = append(intSlice, i)
		fmt.Printf("i = %d; intSlice = %#v; len = %d; cap = %d\n",
			i, intSlice, len(intSlice), cap(intSlice))
	}
	// END OMIT

	fmt.Printf("intSlice = %#v\n", intSlice)
}
