package main

import "fmt"

// START OMIT
type databaseError struct {
	host    string
	message string
}

func (e databaseError) Error() string {
	return fmt.Sprintf("error on host %s: %s", e.host, e.message)
}

// END OMIT
