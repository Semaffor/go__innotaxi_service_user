package model

type Order struct {
	TaxiType TaxiType // тут taxi
	from     Address
	to       Address
}
