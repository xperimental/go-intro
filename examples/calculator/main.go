package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/calculate", calcHandler)
	http.Handle("/", http.FileServer(http.Dir("web/")))

	// Start server
	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
