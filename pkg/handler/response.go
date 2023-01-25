package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type errorStruct struct {
	Message string `json:"message"`
}

func newErrorResponse(ctx *gin.Context, errorCode int, message string) {
	log.Printf("Error message: %s", message)
	ctx.AbortWithStatusJSON(errorCode, errorStruct{message})
}
