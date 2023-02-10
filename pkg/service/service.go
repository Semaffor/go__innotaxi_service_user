package service

import (
	"github.com/Semaffor/go__innotaxi_service_user/pkg/service/log"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/service/token"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/service/user"
)

type Aggregator struct {
	log   *log.Service
	user  *user.Service
	token *token.Service
}

func NewAggregator(
	servicesMongo *log.Service,
	servicesPostgre *user.Service,
	servicesRedis *token.Service,
) *Aggregator {
	return &Aggregator{
		log:   servicesMongo,
		user:  servicesPostgre,
		token: servicesRedis,
	}
}

func (s *Aggregator) GetUserService() *user.Service {
	return s.user
}

func (s *Aggregator) GetLogService() *log.Service {
	return s.log
}

func (s *Aggregator) GetTokenService() *token.Service {
	return s.token
}
