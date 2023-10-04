package controllers

import (
	"github.com/gin-gonic/gin"
	"product_move/internal/domains"
	"product_move/internal/helpers"
	"product_move/internal/infrastructure"
	"product_move/internal/middleware"
	"product_move/internal/repositories"
	"product_move/internal/services"
	"time"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		authService: services.AuthService{
			AuthRep: repositories.AuthRepository{},
		},
	}
}

func (a *AuthController) Build() {
	infrastructure.PostMappingNoAuth("/login", a.LoginHandler)
}

func (a *AuthController) LoginHandler(ctx *gin.Context) {
	data := &domains.AuthRequest{}
	err := ctx.BindJSON(data)
	if err != nil {
		helpers.WriteError(ctx, 400, err)
		return
	}
	login, err := a.authService.Login(data.Username, data.Password)
	if err != nil {
		helpers.WriteError(ctx, 500, err)
		return
	}
	if login {
		token, err := middleware.GenerateAuthToken(data.Username, time.Hour*24)
		if err != nil {
			helpers.WriteError(ctx, 500, err)
			return
		}
		helpers.Write(ctx, 200, gin.H{
			"username": data.Username,
			"token":    token,
		})
	}
	helpers.Write(ctx, 401, "Wrong username or password")
}
