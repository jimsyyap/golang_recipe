package main

import (
    "github.com/jimsyyap/celeritas"
)

type application struct {
    App *celeritas.Celeritas
}

func main() {
    c := initApplication()
    c.App.ListenAndServe()
}
