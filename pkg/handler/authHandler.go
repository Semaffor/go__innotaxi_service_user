package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/domain"
)

func (h *Handler) LogIn(ctx *gin.Context) {
	userCredentials := domain.UserCredentials{}

	if err := ctx.BindJSON(&userCredentials); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "Invalid user data")
		return
	}

	user, err := h.ServicePostgre.User.Authentication(&userCredentials)

	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.ServiceRedis.Authorization.CreateSession(&user)

	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, domain.JwtTokens{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	})
}

func (h *Handler) SignUp(ctx *gin.Context) {
}
