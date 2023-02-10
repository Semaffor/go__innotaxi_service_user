package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	jwtLib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/auth/jwt/model"
)

const (
	authorizationHeader = "Authorization"
	claimId             = "userId"
	claimRole           = "role"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	claims, err := h.checkIsAuth(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(claimId, claims.Id)
	c.Set(claimRole, claims.Role)
}

func (h *Handler) checkIsAuth(c *gin.Context) (*model.JwtClaims, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return nil, errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return nil, errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return nil, errors.New("token is empty")
	}
	claimsMap, err := h.services.GetTokenService().Authorization.GetAuthManager().ParseJwt(headerParts[1])
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
	id, ok := claims[claimId].(int)
	if !ok || id == 0 {
		return nil, fmt.Errorf("incorrect data in field `%s`: %d", claimId, id)
	}

	role, ok := claims[claimRole].(string)
	if !ok || role == "" {
		return nil, fmt.Errorf("incorrect data in field `%s`: %d", claimRole, id)
	}

	return &model.JwtClaims{
		StandardClaims: jwtLib.StandardClaims{},
		UserId:         id,
		Role:           role,
	}, nil
}
