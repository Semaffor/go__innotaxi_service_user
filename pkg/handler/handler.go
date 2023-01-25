package handler

import (
	"github.com/Semaffor/go__innotaxi_service_user/pkg/service/mongodb"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/service/postgres"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	servicesMongo   *mongodb.ServiceMongo
	servicesPostgre *postgres.ServicePostgres
	// tokenManager    auth.TokenManager
}

func NewHandler(servicesMongo *mongodb.ServiceMongo, servicesPostgre *postgres.ServicePostgres) *Handler {
	return &Handler{}
	// return &Handler{servicesMongo: servicesMongo, servicesPostgre: servicesPostgre}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/logIn", h.logIn)
		auth.POST("/signUp", h.signUp)
	}

	return router
}
