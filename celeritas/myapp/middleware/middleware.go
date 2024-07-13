package middleware

import (
    "myapp/data"
    "github.com/jimsyyap/celeritas"
)

type Middleware struct {
    App *celeritas.celeritas
    Models data.Models
}
