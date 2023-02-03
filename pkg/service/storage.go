package service

import (
	"github.com/Semaffor/go__innotaxi_service_user/pkg/service/log"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/service/token"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/service/user"
)

type AggregatorI interface {
	GetUserService() *user.Service
	GetLogService() *log.Service
	GetAuthService() *token.Service
}
