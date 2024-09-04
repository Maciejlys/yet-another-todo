package web

import (
	"github.com/Maciejlys/yet-another-todo/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	*chi.Mux
	store models.Store
}

func NewHandler(store models.Store) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}

	h.Use(middleware.Logger)

	todos := TodoHandler{store: store}

	h.Route("/todos", func(r chi.Router) {
		r.Get("/", todos.List())
		r.Get("/{id}", todos.Get())
		r.Post("/", todos.Create())
		r.Delete("/{id}", todos.Delete())
		r.Patch("/{id}", todos.Edit())
	})

	return h
}
