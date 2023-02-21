package handler

import (
	"errors"
	"fmt"
	"strings"

	jwtLib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/auth/jwt/model"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/errbase"
)

const (
	authorizationHeader = "Authorization"
	ClaimId             = "userId"
	ClaimRole           = "role"
)

func (h *Handler) userIdentity(ctx *gin.Context) {
	claims, err := h.checkIsAuth(ctx)
	if err != nil {
		errbase.NewErrorResponse(ctx, err)
		return
	}

	ctx.Set(ClaimId, claims.UserId)
	ctx.Set(ClaimRole, claims.Role)
}

func (h *Handler) checkIsAuth(ctx *gin.Context) (*model.JwtClaims, error) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		return nil, errbase.InvalidCredentialsError("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return nil, errbase.InvalidCredentialsError("empty auth header")
	}

	if len(headerParts[1]) == 0 {
		return nil, errbase.InvalidCredentialsError("token is empty")
	}
	claimsMap, err := h.services.TokenService().AuthManager().ParseJwt(headerParts[1])
	if err != nil {
		return nil, err
	}

	claims, err := grabFields(claimsMap)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func grabFields(claims map[string]interface{}) (*model.JwtClaims, error) {
	idFloat, ok := claims[ClaimId].(float64)
	if !ok {
		return nil, errors.New("type assertion error")
	}
	id := int(idFloat)
	if id == 0 {
		return nil, errbase.InvalidCredentialsError(fmt.Sprintf("incorrect id in field `%s`: %d", ClaimId, id))
	}

	role, ok := claims[ClaimRole].(string)
	if !ok || role == "" {
		return nil, errbase.InvalidCredentialsError(fmt.Sprintf("incorrect data in field `%s`: %s", ClaimRole, role))
	}

	return &model.JwtClaims{
		StandardClaims: jwtLib.StandardClaims{},
		UserId:         id,
		Role:           role,
	}, nil
}
