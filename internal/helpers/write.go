package helpers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"product_move/internal/exceptions"
)

func Write(ctx *gin.Context, code int, obj interface{}) {
	ctx.JSON(code, gin.H{
		"code":    code,
		"message": http.StatusText(code),
		"result":  obj,
	})
}

func WriteSuccess(ctx *gin.Context, obj interface{}) {
	Write(ctx, http.StatusOK, obj)
}

func WriteErrorWithCode(ctx *gin.Context, code int, err any) {
	switch err.(type) {
	case exceptions.Error:
		ctx.AbortWithStatusJSON(code, gin.H{
			"code":    code,
			"message": http.StatusText(code),
			"result":  err.(exceptions.Error).Error(),
		})
	case error:
		ctx.AbortWithStatusJSON(code, gin.H{
			"code":    code,
			"message": http.StatusText(code),
			"result":  err.(error).Error(),
		})
	default:
		log.Fatal("WriteErrorWithCode: Can only cast from error or exceptions.Errors")
	}
}

func WriteError(ctx *gin.Context, err any) {
	switch err.(type) {
	case exceptions.Error:
		ctx.AbortWithStatusJSON(err.(exceptions.Error).Code(), gin.H{
			"code":    err.(exceptions.Error).Code(),
			"message": http.StatusText(err.(exceptions.Error).Code()),
			"result":  err.(exceptions.Error).Error(),
		})
	case error:
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": http.StatusText(http.StatusInternalServerError),
			"result":  err.(error).Error(),
		})
	default:
		log.Fatal("WriteError: Can only cast from or exceptions.Errors")
	}
}
