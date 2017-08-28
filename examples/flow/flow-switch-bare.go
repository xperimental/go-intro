package main

import (
	"io"
	"log"
	"os"
)

func main() {
	fileName := "file.txt"
	file, err := os.Open(fileName)
	switch {
	case os.IsNotExist(err):
		log.Fatalf("File does not exist: %s", fileName)
	case err != nil:
		log.Fatalf("Error opening file: %s", err)
	default:
	}
	defer file.Close()

	// Show file
	io.Copy(os.Stdout, file)
}