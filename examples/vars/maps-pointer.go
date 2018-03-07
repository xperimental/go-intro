package main

import "fmt"

func main() {
	// START OMIT
	m1 := map[string]string{"key1": "value1", "key2": "value2"}
	m2 := m1
	m2["key1"] = "value3" // HL
	// END OMIT
	fmt.Printf("m1 = %#v\n", m1)
	fmt.Printf("m2 = %#v\n", m2)
}
