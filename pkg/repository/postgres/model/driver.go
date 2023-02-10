package model

import "github.com/guregu/null"

const (
	BUSY = 1
	FREE = 2
)

type Driver struct {
	Id         int        `db:"id"`
	UserId     int        `db:"user_id"`
	Status     int        `db:"status"`
	TaxiTypeId int        `db:"taxi_type_id"`
	TotalMark  null.Float `db:"total_mark"`
}
