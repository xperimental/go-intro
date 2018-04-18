package main

// START OMIT
// see io package in stdlib
type Reader interface {
	Read(p []byte) (int, error)
}

// END OMIT
