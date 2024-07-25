package main

import (
	"fmt"
	"github.com/gorilla/context"
	"github.com/jimsyyap/negroni"
	"log"
	"net/http"
)

func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token := r.Header.Get("X-AppToken")
	if token == "bXlfdG9rZW4=" {
		log.Printf("Authorized")
		context.Set(r, "user", "jim")
		next(w, r)
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	user := context.Get(r, "user")
	fmt.Fprintf(w, "Hello %v", user)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(Authorize))
	n.UseHandler(mux)
	n.Run(":8080")
}
