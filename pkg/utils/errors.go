package utils

import (
	"github.com/gin-gonic/gin"
)

func HandleHttpError(ctx *gin.Context, err error, code int) {
	if err == nil {
		return
	}

	ctx.JSON(code, gin.H{
		"error": err.Error(),
	})
}
