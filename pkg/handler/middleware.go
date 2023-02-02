package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	id, role, err := h.checkIsAuth(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set("userId", id)
	c.Set("role", role)
}

func (h *Handler) checkIsAuth(c *gin.Context) (int, string, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return -1, "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return -1, "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return -1, "", errors.New("token is empty")
	}
	claims, err := h.servicesRedis.Authorization.GetAuthManager().ParseJwt(headerParts[1])
	if err != nil {
		return -1, "", errors.New(err.Error())
	}

	return claims["id"].(int), claims["userId"].(string), nil
}
