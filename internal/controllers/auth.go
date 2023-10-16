package controllers

import (
	"github.com/gin-gonic/gin"
	"product_move/internal/domains"
	"product_move/internal/helpers"
	"product_move/internal/infrastructure"
	"product_move/internal/services"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{authService: services.NewAuthService()}
}

func (a *AuthController) Build() {
	infrastructure.PostMappingNoAuth("/login", a.LoginHandler)
}

func (a *AuthController) LoginHandler(ctx *gin.Context) {
	data := &domains.AuthRequest{}
	err := ctx.BindJSON(data)
	if err != nil {
		helpers.WriteError(ctx, err)
		return
	}
	resp, err := a.authService.Login(*data)
	if err != nil {
		helpers.WriteError(ctx, err)
		return
	}
	helpers.WriteSuccess(ctx, resp)
}
