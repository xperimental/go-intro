package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS
	// START OMIT
	fmt.Print("Go runs on ")
	switch os { // os is a string
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
	// END OMIT
}
