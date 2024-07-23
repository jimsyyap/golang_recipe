package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jimsyyap/negroni"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":8080")
}
