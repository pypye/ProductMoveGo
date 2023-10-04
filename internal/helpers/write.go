package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Write(ctx *gin.Context, code int, obj interface{}) {
	ctx.JSON(code, gin.H{
		"code":    code,
		"message": http.StatusText(code),
		"result":  obj,
	})
}

func WriteError(ctx *gin.Context, code int, err error) {
	ctx.AbortWithStatusJSON(code, gin.H{
		"code":    code,
		"message": http.StatusText(code),
		"result":  err.Error(),
	})
}
