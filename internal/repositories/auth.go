package repositories

import (
	"database/sql"
	"product_move/internal/infrastructure"
)

type AuthInterface interface {
	Login(username string, password string) (bool, error)
}

type AuthRepository struct {
}

func (a *AuthRepository) Login(username string, password string) (bool, error) {
	rows, err := infrastructure.GetDB().Query("SELECT * FROM auth WHERE username = ? AND password = ?", username, password)
	if err != nil {
		return false, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	if rows.Next() {
		return true, nil
	}
	return false, nil
}
