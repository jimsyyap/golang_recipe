package handlers

import (
    "net/http"
    "github.com/jimsyyap/celeritas"
)

type Handlers struct {
    App *celeritas.Celeritas
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
    err := h.App.Render(w, r, "home", nil, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
