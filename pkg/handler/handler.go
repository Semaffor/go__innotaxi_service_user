package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/service"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/user/v1")
	{
		auth.POST("/login", h.logIn)
		auth.POST("/signup", h.signUp)
		auth.POST("/auth/refresh", h.userRefresh)

		api := auth.Group("", h.userIdentity)
		{
			api.POST("/logout", h.logout)
			api.PATCH("/", h.updateData)
			api.DELETE("/", h.deleteUser)
		}
	}

	return router
}
