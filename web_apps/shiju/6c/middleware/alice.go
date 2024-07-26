package main

import (
	"github.com/gorilla/handlers"
	"github.com/jimsyyap/alice"
	"io"
	"log"
	"net/http"
	"os"
)

func loggingHandler(next http.Handler) http.Handler {
	logFile, err := os.Openfile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatal(err)
	}
	return handlers.LoggingHandler(logFile, next)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(
		w,
		`<doctype html>
        <html>
        <head>
        <title>Hello World</title>
        </head>
        <body>
        Hello World!
        </body>
        </html>`,
	)
}
