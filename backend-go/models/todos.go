package models

type Todo struct {
	ID   int    `db:"id" json:"id"`
	Task string `db:"task" json:"task"`
	Done string `db:"done" json:"done"`
}

type TodoStore interface {
	Todo(id int) (Todo, error)
	Todos() ([]Todo, error)
	CreateTodo(t *Todo) error
	UpdateTodo(t *Todo, id int) error
	DeleteTodo(id int) error
}
