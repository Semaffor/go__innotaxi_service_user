package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/auth/jwt/model"
)

func (h *Handler) logIn(ctx *gin.Context) {
	userCredentials := model.UserCredentials{}
	if err := ctx.BindJSON(&userCredentials); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "Invalid user data")

		return
	}

	user, err := h.services.GetUserService().Authentication(&userCredentials)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	res, err := h.services.GetTokenService().Authorization.CreateSession(&user)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	ctx.JSON(http.StatusOK, model.JwtTokens{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	})
}

func (h *Handler) signUp(ctx *gin.Context) {
}
