package model

const (
	PREMIUM = "Премиум"
	COMFORT = "Комфорт"
	ECONOMY = "Эконом"
)

type TaxiType struct {
	Id            int    `db:"id"`
	TaxiTypeTitle string `db:"type"`
	Description   string `db:"description"`
}
