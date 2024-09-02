package db

import (
	"fmt"

	"github.com/Maciejlys/yet-another-todo"
	"github.com/jmoiron/sqlx"
)

type TodoStore struct {
	*sqlx.DB
}

func (s *TodoStore) Todo(id int) (todo.Todo, error) {
	var t todo.Todo
	if err := s.Get(&t, `SELECT * FROM todos WHERE id = $1`, id); err != nil {
		return todo.Todo{}, fmt.Errorf("error getting todo: %w", err)
	}
	return t, nil
}

func (s *TodoStore) Todos() ([]todo.Todo, error) {
	var tt []todo.Todo
	if err := s.Select(&tt, `SELECT * FROM todos`); err != nil {
		return []todo.Todo{}, fmt.Errorf("error getting todos: %w", err)
	}
	return tt, nil
}

func (s *TodoStore) CreateTodo(t *todo.Todo) error {
	if err := s.Get(t, `INSERT INTO todos(task, done) VALUES ($1, $2) RETURNING *`,
		t.Task,
		t.Done); err != nil {
		return fmt.Errorf("error creating todo: %w", err)
	}
	return nil
}

func (s *TodoStore) UpdateTodo(t *todo.Todo, id int) error {
	if err := s.Get(t, `UPDATE todos SET task = $1, done = $2 WHERE id = $3 RETURNING *`,
		t.Task,
		t.Done,
		id); err != nil {
		return fmt.Errorf("error updating thread: %w", err)
	}
	return nil
}

func (s *TodoStore) DeleteTodo(id int) error {
	if _, err := s.Exec(`DELETE FROM todos WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting thread: %w", err)
	}
	return nil
}
