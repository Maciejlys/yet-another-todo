package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Maciejlys/yet-another-todo"
	"github.com/Maciejlys/yet-another-todo/utils"
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
			utils.WriteError(w, err, http.StatusBadRequest)
			return
		}

		tt, err := h.store.Todo(id)

		if err != nil {
			utils.WriteError(w, err, http.StatusInternalServerError)
			return
		}

		utils.WriteJSON(w, tt, http.StatusOK)
	}
}

func (h *TodoHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tt, err := h.store.Todos()
		if err != nil {
			utils.WriteError(w, err, http.StatusInternalServerError)
			return
		}

		utils.WriteJSON(w, tt, http.StatusOK)
	}
}

func (h *TodoHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		form := CreateTodoForm{
			Task: r.FormValue("task"),
			Done: r.FormValue("done"),
		}

		if !form.Validate() {
			utils.WriteErrors(w, form.Errors, http.StatusBadRequest)
			return
		}

		if err := h.store.CreateTodo(&todo.Todo{
			Task: form.Task,
			Done: form.Done,
		}); err != nil {
			utils.WriteError(w, err, http.StatusInternalServerError)
			return
		}

		utils.WriteMsg(w, "Created", http.StatusCreated)
	}
}

func (h *TodoHandler) Edit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			utils.WriteError(w, err, http.StatusInternalServerError)
			return
		}

		form := CreateTodoForm{
			Task: r.FormValue("task"),
			Done: r.FormValue("done"),
		}

		if !form.Validate() {
			utils.WriteErrors(w, form.Errors, http.StatusBadRequest)
			return
		}

		if err := h.store.UpdateTodo(&todo.Todo{
			Task: form.Task,
			Done: form.Done,
		}, id); err != nil {
			utils.WriteError(w, err, http.StatusInternalServerError)
			return
		}

		utils.WriteMsg(w, "Updated", http.StatusOK)
	}
}

func (h *TodoHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			utils.WriteError(w, err, http.StatusInternalServerError)
			return
		}

		err = h.store.DeleteTodo(id)

		if err != nil {
			utils.WriteError(w, err, http.StatusInternalServerError)
			return
		}

		utils.WriteMsg(w, "Removed", http.StatusOK)
	}
}
