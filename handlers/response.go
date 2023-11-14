package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, gin.H{
		"error": message,
	})
}

func SuccesResponse(ctx *gin.Context, status int, message string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}
