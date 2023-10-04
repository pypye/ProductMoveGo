package infrastructure

import (
	"database/sql"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("mysql", "root:password@(localhost:3306)/product_move")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) Query(query string, args ...any) (*sql.Rows, error) {
	return d.db.Query(query, args...)
}
