package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/service"
)

type Handler struct {
	services *service.Aggregator
}

func NewHandler(services *service.Aggregator) *Handler {
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
	}

	return router
}
