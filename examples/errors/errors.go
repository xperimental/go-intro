package main

import "errors"

// START OMIT
func add(a, b int) int {
	return a + b
}

func divide(a, b int) (int, error) { // HL
	if b == 0 {
		return 0, errors.New("divisor is zero")
	}

	return a / b, nil
}

// END OMIT
