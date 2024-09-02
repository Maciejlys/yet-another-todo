package todo

type Todo struct {
	ID   int    `db:"id"`
	Task string `db:"task"`
	Done bool   `db:"done"`
}

type TodoStore interface {
	Todo(id int) (Todo, error)
	Todos() ([]Todo, error)
	CreateTodo(t *Todo) error
	// UpdateTodo(t *Todo) error
	// DeleteTodo(id int) error
}

type Store interface {
	TodoStore
}
