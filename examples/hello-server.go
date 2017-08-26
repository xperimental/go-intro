package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// do not ignore errors in your code!
	me, _ := os.Hostname()

	http.HandleFunc("/", handler)
	log.Printf("%s is listening on port 8080...", me)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	remote := r.RemoteAddr

	w.WriteHeader(http.StatusOK)
	log.Printf("Hello to %s!", remote)
	fmt.Fprintf(w, "Hello %s!", remote)
}
