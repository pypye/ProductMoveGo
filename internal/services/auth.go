package services

import (
	"product_move/internal/domains"
	"product_move/internal/exceptions"
	"product_move/internal/middleware"
	"product_move/internal/repositories"
	"time"
)

type AuthService struct {
	AuthRep repositories.AuthRepository
}

func (a *AuthService) Login(request domains.AuthRequest) (*domains.AuthResponse, error) {
	loginValid, err := a.AuthRep.Login(request)
	if err != nil {
		return nil, err
	}
	if loginValid {
		token, err := middleware.GenerateAuthToken(request.Username, time.Hour*24)
		if err != nil {
			return nil, err
		}
		return &domains.AuthResponse{
			Username: request.Username,
			Token:    token,
		}, nil
	}
	return nil, &exceptions.WrongIdentityError{}
}
