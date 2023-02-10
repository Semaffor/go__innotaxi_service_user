package service

import (
	"github.com/Semaffor/go__innotaxi_service_user/pkg/service/log"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/service/token"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/service/user"
)

type AggregatorI interface {
	UserService() *user.Service
	LogService() *log.Service
	AuthService() *token.Service
}
