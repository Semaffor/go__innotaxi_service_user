package mongo

import "github.com/jmoiron/sqlx"

type Logger interface{}

type LogsRepo struct {
	Logger
}

func NewLogsRepository(db *sqlx.DB) *LogsRepo {
	return &LogsRepo{}
}
