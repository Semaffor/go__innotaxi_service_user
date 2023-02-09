package model

import "github.com/guregu/null"

type Feedback struct {
	Id                 int         `db:"id"`
	DriverId           int         `db:"driver_id"`
	CustomerId         int         `db:"customer_id"`
	MarkFromUser       null.Float  `db:"mark_from_user"`
	MarkFromDriver     null.Float  `db:"mark_from_driver"`
	FeedbackFromUser   null.String `db:"feedback_from_user"`
	FeedbackFromDriver null.String `db:"feedback_from_driver"`
	OrderId            int         `db:"order_id"`
}
