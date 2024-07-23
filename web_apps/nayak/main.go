package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	PORT = "8080"
)

func serveDynamic(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Time: %s", time.Now())
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static.html")
}

func main() {
	http.HandleFunc("/static", serveStatic)
	http.HandleFunc("/", serveDynamic)
	http.ListenAndServe(":"+PORT, nil)
}
