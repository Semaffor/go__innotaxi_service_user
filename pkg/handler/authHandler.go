package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/auth/jwt/model"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/errbase"
	form "github.com/Semaffor/go__innotaxi_service_user/pkg/handler/model"
)

func (h *Handler) logIn(ctx *gin.Context) {
	var userCredentials form.UserLoginInput
	if err := ctx.BindJSON(&userCredentials); err != nil {
		errbase.NewErrorResponse(ctx, errbase.InvalidInput(err))
		return
	}

	user, err := h.services.UserService().Authenticate(ctx, &userCredentials)
	if err != nil {
		errbase.NewErrorResponse(ctx, err)
		return
	}

	res, err := h.services.TokenService().CreateSession(ctx, user.Id, user.Role)
	if err != nil {
		errbase.NewErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, errbase.NewJSONSuccessResponse(
		model.JwtTokens{
			AccessToken:  res.AccessToken,
			RefreshToken: res.RefreshToken,
		}),
	)
}

func (h *Handler) signUp(ctx *gin.Context) {
	var user form.UserRegistrationInput
	if err := ctx.ShouldBindJSON(&user); err != nil {
		errbase.NewErrorResponse(ctx, errbase.InvalidInput(err))
		return
	}

	err := h.services.UserService().Register(ctx, &user)
	if err != nil {
		errbase.NewErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, errbase.NewJSONSuccessResponse(nil))
}

func (h *Handler) userRefresh(ctx *gin.Context) {
	var input model.RefreshInput
	if err := ctx.BindJSON(&input); err != nil {
		errbase.NewErrorResponse(ctx, errbase.InvalidInput(err))
		return
	}

	res, err := h.services.TokenService().RefreshTokens(ctx.Request.Context(), input.Token)
	if err != nil {
		errbase.NewErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, errbase.NewJSONSuccessResponse(
		model.JwtTokens{
			AccessToken:  res.AccessToken,
			RefreshToken: res.RefreshToken,
		}),
	)
}

func (h *Handler) logout(ctx *gin.Context) {
	var input model.RefreshInput
	if err := ctx.BindJSON(&input); err != nil {
		errbase.NewErrorResponse(ctx, errbase.InvalidInput(err))
		return
	}

	userId := ctx.GetInt(claimId)
	err := h.services.TokenService().LogoutSingle(ctx, userId, input.Token)
	if err != nil {
		errbase.NewErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, errbase.NewJSONSuccessResponse(nil))
}
