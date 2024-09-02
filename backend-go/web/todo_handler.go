package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Maciejlys/yet-another-todo"
	"github.com/go-chi/chi/v5"
)

type TodoHandler struct {
	store todo.Store
}

func (h *TodoHandler) Get() http.HandlerFunc {
	type data struct {
		Todos []todo.Todo
	}

	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tt, err := h.store.Todo(id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(tt)
	}
}

func (h *TodoHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tt, err := h.store.Todos()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(tt)
	}
}

func (h *TodoHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		form := CreateTodoForm{
			Task: r.FormValue("task"),
			Done: r.FormValue("done"),
		}

		if !form.Validate() {
			http.Error(w, "Form is not valid", http.StatusBadRequest)
			return
		}

		if err := h.store.CreateTodo(&todo.Todo{
			Task: form.Task,
			Done: form.Done,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, "Created")
	}
}
