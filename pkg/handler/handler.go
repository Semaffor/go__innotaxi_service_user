package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/service/mongodb"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/service/postgres"
)

type Handler struct {
	servicesMongo   *mongodb.ServiceMongo
	servicesPostgre *postgres.ServicePostgres
}

func NewHandler(serviceMongo *mongodb.ServiceMongo, servicePostgre *postgres.ServicePostgres) *Handler {
	fmt.Println(serviceMongo, servicePostgre)
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	v1 := router.Group("/api/v1")

	auth := v1.Group("/auth")
	{
		auth.POST("/logIn", h.logIn)
		auth.POST("/signUp", h.signUp)
	}

	return router
}
