package store

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Store ...
type Store struct {
	config *Config
	db     *sql.DB
}

//  New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	db, err := sql.Open("sqlite3", "store.db")
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// Close ...
func (s *Store) Close() {
	s.db.Close()
}
