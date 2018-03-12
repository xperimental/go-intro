package main

import (
	"errors"
	"fmt"
)

// START OMIT
func createHello(name string) (string, error) { // HL
	if len(name) == 0 {
		return "", errors.New("need a name")
	}

	hello := fmt.Sprintf("Hello %s!", name)
	return hello, nil
}

func main() {
	hello, err := createHello("") // HL
	if err != nil {
		fmt.Printf("Error creating hello: %s\n", err)
		return
	}

	fmt.Println(hello)
}

// END OMIT
