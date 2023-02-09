package model

type TaxiType struct {
	Id            int    `db:"id"`
	TaxiTypeValue int    `db:"type"`
	Description   string `db:"description"`
}
