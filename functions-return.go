package main

import "fmt"

// START OMIT
func createHello(name string) string { // HL
	return fmt.Sprintf("Hello %s!", name)
}

func main() {
	fmt.Println(createHello("go-intro"))
}

// END OMIT
