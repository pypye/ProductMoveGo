package db_handler

import (
	"database/sql"
)

type MySqlStore struct {
	db *sql.DB
}

func NewMySqlStore() (*MySqlStore, error) {
	db, err := sql.Open("mysql", "root:password@(localhost:3306)/product_move")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &MySqlStore{db: db}, nil
}

func (mySqlStore *MySqlStore) GetDB() *sql.DB {
	return mySqlStore.db
}
