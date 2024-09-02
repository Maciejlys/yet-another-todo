package web

import (
	"fmt"
	"net/http"

	todo "github.com/Maciejlys/yet-another-todo"
)

type TodoHandler struct {
	store todo.Store
}

func (h *TodoHandler) List() http.HandlerFunc {
	type data struct {
		Todos []todo.Todo
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tt, err := h.store.Todos()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, tt)
	}
}

func (h *TodoHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		form := CreateTodoForm{
			Task: r.FormValue("task"),
			Done: r.FormValue("done"),
		}

		if !form.Validate() {
			http.Redirect(w, r, r.Referer(), http.StatusFound)
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
