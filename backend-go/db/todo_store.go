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
		return todo.Todo{}, fmt.Errorf("error getting thread: %w", err)
	}
	return t, nil
}

func (s *TodoStore) Todos() ([]todo.Todo, error) {
	var tt []todo.Todo
	if err := s.Select(&tt, `SELECT * FROM todos`); err != nil {
		return []todo.Todo{}, fmt.Errorf("error getting threads: %w", err)
	}
	return tt, nil
}

func (s *TodoStore) CreateTodo(t *todo.Todo) error {
	if err := s.Get(t, `INSERT INTO todos VALUES ($1, $2) RETURNING *`,
		t.Task,
		t.Done); err != nil {
		return fmt.Errorf("error creating thread: %w", err)
	}
	return nil
}

// func (s *ThreadStore) UpdateThread(t *goreddit.Thread) error {
// 	if err := s.Get(t, `UPDATE threads SET title = $1, description = $2 WHERE id = $3 RETURNING *`,
// 		t.Title,
// 		t.Description,
// 		t.ID); err != nil {
// 		return fmt.Errorf("error updating thread: %w", err)
// 	}
// 	return nil
// }
//
// func (s *ThreadStore) DeleteThread(id uuid.UUID) error {
// 	if _, err := s.Exec(`DELETE FROM threads WHERE id = $1`, id); err != nil {
// 		return fmt.Errorf("error deleting thread: %w", err)
// 	}
// 	return nil
// }
