package model

type TaxiStatus int64

const (
	Free TaxiStatus = iota
	Booked
	Offline
)

type Taxi struct {
	TaxiType TaxiType
	Status   TaxiStatus
}
