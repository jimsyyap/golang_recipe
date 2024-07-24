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
}
