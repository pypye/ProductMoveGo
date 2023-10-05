package repositories

import (
	"errors"
	"gorm.io/gorm"
	"product_move/internal/domains"
	"product_move/internal/exceptions"
	"product_move/internal/infrastructure"
)

type AuthInterface interface {
	Login(username string, password string) (bool, error)
}

type AuthRepository struct {
}

func (a *AuthRepository) Login(request domains.AuthRequest) (bool, error) {
	db := infrastructure.GetDB().Get()
	err := db.Where("username = ? AND password = ?", request.Username, request.Password).First(&request).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, &exceptions.WrongIdentityError{}
		}
		return false, err
	}
	return true, nil
}
