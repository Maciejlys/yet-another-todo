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

		w.Header().Set("Content-Type", "application/json")
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

		w.Header().Set("Content-Type", "application/json")
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
			http.Error(w, form.Errors.String(), http.StatusBadRequest)
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

func (h *TodoHandler) Edit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		form := CreateTodoForm{
			Task: r.FormValue("task"),
			Done: r.FormValue("done"),
		}

		if !form.Validate() {
			http.Error(w, form.Errors.String(), http.StatusBadRequest)
			return
		}

		if err := h.store.UpdateTodo(&todo.Todo{
			Task: form.Task,
			Done: form.Done,
		}, id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, "Updated")
	}
}

func (h *TodoHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = h.store.DeleteTodo(id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, "Removed")
	}
}
