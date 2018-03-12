package main

func main() {
	// START OMIT
	var fiveIntArray [5]int // five elements
	var tenIntArray [10]int // ten elements -> different type!

	// won't compile
	fiveIntArray = tenIntArray
	// END OMIT
}
