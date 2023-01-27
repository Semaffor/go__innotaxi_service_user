package go__innotaxi_service_user

import (
	"context"
	"net/http"
	"time"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/config"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(serverConfig *config.Server, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + serverConfig.Port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
