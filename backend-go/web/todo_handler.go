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
