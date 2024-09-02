package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	*TodoStore
}

func migrate(db *sqlx.DB) {
	const migration = `CREATE TABLE IF NOT EXISTS todos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        task TEXT NOT NULL,
        done INTEGER NOT NULL
    )`

	db.Exec(migration)
}

func NewStore() (*Store, error) {
	db, err := sqlx.Open("sqlite3", "todos.db")

	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	migrate(db)

	return &Store{
		TodoStore: &TodoStore{DB: db},
	}, nil
}
