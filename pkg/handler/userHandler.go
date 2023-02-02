package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"test": "ok",
	})
}
