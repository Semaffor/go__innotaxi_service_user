package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) logIn(ctx *gin.Context) {
	ctx.JSON(200, map[string]interface{}{
		"payload": "example",
	})
}

func (h *Handler) signUp(ctx *gin.Context) {

}
