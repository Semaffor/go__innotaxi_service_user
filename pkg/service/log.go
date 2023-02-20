package service

import (
	"context"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/repository"
)

type Logs struct {
	repo repository.Logs
}

func NewLogsService(repo repository.Logs) *Logs {
	return &Logs{repo: repo}
}

func (l *Logs) WriteLog(ctx context.Context, log string) error {
	return nil
}
