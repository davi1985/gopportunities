package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content_type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(
	ctx *gin.Context,
	code int,
	operation string,
	data interface{},
) {
	ctx.Header("Content_type", "application/json")
	ctx.JSON(code, gin.H{
		"message": fmt.Sprintf(
			"operation from handler: %s successful",
			operation,
		),
		"data": data,
	})
}
