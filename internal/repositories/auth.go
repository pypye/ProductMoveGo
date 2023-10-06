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

func matchingIdentity(request domains.AuthRequest) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("username = ? AND password = ?", request.Username, request.Password)
	}
}

func (a *AuthRepository) Login(request domains.AuthRequest) (bool, error) {
	err := infrastructure.GetDB().Scopes(matchingIdentity(request)).First(&request).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, &exceptions.WrongIdentityError{}
		}
		return false, err
	}
	return true, nil
}
