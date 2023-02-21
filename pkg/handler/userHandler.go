package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/errbase"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/handler/model"
)

func (h *Handler) updateData(ctx *gin.Context) {
	var user model.UserUpdateInput
	if err := ctx.BindJSON(&user); err != nil {
		errbase.NewErrorResponse(ctx, errbase.InvalidInput(err))
		return
	}

	user.Id = ctx.GetInt(ClaimId)
	err := h.services.UserService().UpdateUser(ctx, &user)
	if err != nil {
		errbase.NewErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, errbase.NewJSONSuccessResponse(nil))
}

func (h *Handler) deleteUser(ctx *gin.Context) {
	userId := ctx.GetInt(ClaimId)
	err := h.services.UserService().DeleteUser(ctx, userId)
	if err != nil {
		errbase.NewErrorResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, errbase.NewJSONSuccessResponse(nil))
}
