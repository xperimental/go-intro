package main

import (
	"errors"
	"fmt"
)

func sayHello(name string) error {
	if len(name) == 0 {
		return errors.New("name can not be empty")
	}

	fmt.Printf("Hello %s!\n", name)
	return nil
}

func main() {
	// START OMIT
	name := ""

	// func sayHello(name string) error
	if err := sayHello(name); err != nil { // HL
		fmt.Printf("Error saying hello: %s\n", err)
	}

	// err is not defined here // HL
	// fmt.Println(err)
	// END OMIT
}
