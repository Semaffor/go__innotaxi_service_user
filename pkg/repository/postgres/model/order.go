package model

import (
	"time"
)

const (
	InProgress = 1
	Finished   = 2
)

type Order struct {
	id         int       `db:"id"`
	from       string    `db:"from"`
	to         string    `db:"to"`
	date       time.Time `db:"date"`
	status     int       `db:"status"`
	driverId   int       `db:"driver_id"`
	taxiTypeId int       `db:"taxi_type_id"`
	userId     int       `db:"user_id"`
}
