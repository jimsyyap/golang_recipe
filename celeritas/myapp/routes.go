package main

import (
    "net/http"
    "github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {
    a.App.Routes.Get("/", a.Handlers.Home)

    //to add contacts page, do
    //a.App.Routes.Get("/contacts", a.Handlers.Contacts)

    fileServer := http.FileServer(http.Dir("./public"))
    a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

    return a.App.Routes
}
