package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) logIn(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"payload": "example",
	})
}

func (h *Handler) signUp(ctx *gin.Context) {

}
