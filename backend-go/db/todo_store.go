package db

import (
	"fmt"

	"github.com/Maciejlys/yet-another-todo/models"
	"github.com/jmoiron/sqlx"
)

type TodoStore struct {
	*sqlx.DB
}

func (s *TodoStore) Todo(id int) (models.Todo, error) {
	var t models.Todo
	if err := s.Get(&t, `SELECT * FROM todos WHERE id = $1`, id); err != nil {
		return models.Todo{}, fmt.Errorf("error getting todo: %w", err)
	}
	return t, nil
}

func (s *TodoStore) Todos() ([]models.Todo, error) {
	var tt []models.Todo
	if err := s.Select(&tt, `SELECT * FROM todos`); err != nil {
		return []models.Todo{}, fmt.Errorf("error getting todos: %w", err)
	}
	return tt, nil
}

func (s *TodoStore) CreateTodo(t *models.Todo) error {
	if err := s.Get(t, `INSERT INTO todos(task, done) VALUES ($1, $2) RETURNING *`,
		t.Task,
		t.Done); err != nil {
		return fmt.Errorf("error creating todo: %w", err)
	}
	return nil
}

func (s *TodoStore) UpdateTodo(t *models.Todo, id int) error {
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
