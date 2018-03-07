package main

import "fmt"

func main() {
	// START OMIT
	var stringMap = make(map[string]string)
	stringMap["key1"] = "value1"
	var stringMap2 = map[int]string{} // alternative to make()
	// END OMIT
	fmt.Printf("stringMap = %#v\n", stringMap)
	fmt.Printf("stringMap2 = %#v\n", stringMap2)
}
