package handler

import (
	"github.com/Semaffor/go__innotaxi_service_user/pkg/service/mongodb"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/service/postgres"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/service/redis"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	serviceMongo   *mongodb.ServiceMongo
	servicePostgre *postgres.ServicePostgres
	serviceRedis   *redis.ServiceRedis
}

func NewHandler(
	serviceMongo *mongodb.ServiceMongo,
	servicePostgre *postgres.ServicePostgres,
	serviceRedis *redis.ServiceRedis,
) *Handler {
	return &Handler{
		serviceMongo,
		servicePostgre,
		serviceRedis,
	}
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
