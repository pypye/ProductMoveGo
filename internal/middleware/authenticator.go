package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"product_move/internal/errors"
	"product_move/internal/helpers"
	"strings"
	"time"
)

func GenerateToken(obj jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, obj)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateAuthToken(username string, duration time.Duration) (string, error) {
	token := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(duration).Unix(),
	}
	return GenerateToken(token)
}

func Authenticate(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		helpers.WriteError(ctx, 401, &errors.UnauthorizedError{})
		return
	}
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		helpers.WriteError(ctx, 401, err)
		return
	}
	if !token.Valid {
		helpers.WriteError(ctx, 401, &errors.UnauthorizedError{})
		return
	}
	ctx.Next()
}
