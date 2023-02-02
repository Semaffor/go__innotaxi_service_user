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
	// servicesRedis   *redis.ServiceRedis
}

func NewHandler(serviceMongo *mongodb.ServiceMongo, servicePostgre *postgres.ServicePostgres) *Handler {
	fmt.Println(serviceMongo, servicePostgre)
	return &Handler{}
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
