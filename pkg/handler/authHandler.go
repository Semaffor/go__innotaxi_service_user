package handler

import (
	"github.com/Semaffor/go__innotaxi_service_user/pkg/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) logIn(ctx *gin.Context) {
	userCredentials := domain.UserCredentials{}

	if err := ctx.BindJSON(&userCredentials); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "Invalid user data")
		return
	}

	user, err := h.servicePostgre.User.Authentication(&userCredentials)

	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.serviceRedis.Authorization.CreateSession(user)

	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, domain.JwtTokens{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	})
}

func (h *Handler) signUp(ctx *gin.Context) {
}
