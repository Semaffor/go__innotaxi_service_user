package service

import (
	"github.com/Semaffor/go__innotaxi_service_user/pkg/service/log"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/service/token"
	"github.com/Semaffor/go__innotaxi_service_user/pkg/service/user"
)

type Aggregator struct {
	servicesMongo   *log.Service
	servicesPostgre *user.Service
	servicesRedis   *token.Service
}

func NewAggregator(
	servicesMongo *log.Service,
	servicesPostgre *user.Service,
	servicesRedis *token.Service,
) *Aggregator {
	return &Aggregator{
		servicesMongo:   servicesMongo,
		servicesPostgre: servicesPostgre,
		servicesRedis:   servicesRedis,
	}
}

func (s *Aggregator) GetUserService() *user.Service {
	return s.servicesPostgre
}

func (s *Aggregator) GetLogService() *log.Service {
	return s.servicesMongo
}

func (s *Aggregator) GetTokenService() *token.Service {
	return s.servicesRedis
}
