package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/auth"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/service/mongodb"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/service/postgres"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/service/redis"
)

type Handler struct {
	ServiceMongo   *mongodb.ServiceMongo
	ServicePostgre *postgres.ServicePostgres
	ServiceRedis   *redis.ServiceRedis
	AuthManager    *auth.Manager
}

func NewHandler(
	serviceMongo *mongodb.ServiceMongo,
	servicePostgre *postgres.ServicePostgres,
	serviceRedis *redis.ServiceRedis,
	authManager *auth.Manager,
) *Handler {
	return &Handler{
		serviceMongo,
		servicePostgre,
		serviceRedis,
		authManager,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	v1 := router.Group("/api/v1")

	v1.POST("/logIn", h.LogIn)
	v1.POST("/signUp", h.SignUp)

	authed := v1.Group("/", h.UserIdentity)
	{
		authed.GET("/users", h.GetAll)
	}

	return router
}
