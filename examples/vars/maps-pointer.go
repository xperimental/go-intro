package main

import "fmt"

func main() {
	// START OMIT
	m1 := map[string]string{"key": "value1"}
	m2 := m1              // m2 is the same map
	m2["key2"] = "value2" // this modifies both maps // HL
	// END OMIT
	fmt.Printf("m1 = %#v\n", m1)
	fmt.Printf("m2 = %#v\n", m2)
}
