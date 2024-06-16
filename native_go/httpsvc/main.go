package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8111
	http.HandleFunc("/", systemdServer)
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func systemdServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
