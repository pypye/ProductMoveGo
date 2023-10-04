package services

import "product_move/internal/repositories"

type AuthService struct {
	AuthRep repositories.AuthRepository
}

func (a *AuthService) Login(username string, password string) (bool, error) {
	return a.AuthRep.Login(username, password)
}
