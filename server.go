package go__innotaxi_service_user

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/config"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(serverConfig *config.ServerConfig, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port),
		Handler:        handler,
		MaxHeaderBytes: 1 << serverConfig.MaxHeaderBytes, // 1 MB
		ReadTimeout:    serverConfig.ReadTimeout * time.Second,
		WriteTimeout:   serverConfig.WriteTimeout * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
