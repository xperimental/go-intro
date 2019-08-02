package main

import "fmt"

func main() {
	// START OMIT
	var stringMap map[string]string // == nil
	fmt.Printf("stringMap = %#v\n", stringMap)
	stringMap["key1"] = "value1" // crash here // HL
	// END OMIT
}
