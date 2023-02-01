package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

type errorStruct struct {
	Message string `json:"message"`
}

func NewErrorResponse(ctx *gin.Context, errorCode int, message string) {
	log.Printf("Error message: %s", message)
	ctx.AbortWithStatusJSON(errorCode, errorStruct{message})
}
