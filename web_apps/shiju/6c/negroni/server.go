package main

import (
	"fmt"
	"github.com/jimsyyap/negroni"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func main() {
	mux := http.NewServerMux()
	mux.HandleFunc("/", index)
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":8080")
}
