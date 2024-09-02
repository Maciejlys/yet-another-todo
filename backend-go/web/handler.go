package web

import (
	"github.com/Maciejlys/yet-another-todo"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	*chi.Mux
	store todo.Store
}

func NewHandler(store todo.Store) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}

	todos := TodoHandler{store: store}

	h.Route("/todos", func(r chi.Router) {
		r.Get("/", todos.List())
	})

	return h
}
