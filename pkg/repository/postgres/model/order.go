package model

import (
	"time"
)

const (
	InProgress = 1
	Finished   = 2
)

type Order struct {
	Id         int       `db:"id"`
	From       string    `db:"from"`
	To         string    `db:"to"`
	Date       time.Time `db:"date"`
	Status     int       `db:"status"`
	DriverId   int       `db:"driver_id"`
	TaxiTypeId int       `db:"taxi_type_id"`
	UserId     int       `db:"user_id"`
}
