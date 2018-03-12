package main

// START OMIT
// taken from io package
type Reader interface {
	Read(p []byte) (int, error)
}

// END OMIT
