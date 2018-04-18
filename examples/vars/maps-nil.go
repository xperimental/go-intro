package main

func main() {
	// START OMIT
	var stringMap map[string]string // = nil
	stringMap["key1"] = "value1"    // crash here // HL
	// END OMIT
}
