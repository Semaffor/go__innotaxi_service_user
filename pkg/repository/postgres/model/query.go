package model

type SearchArg struct {
	TableName string `json:"-"`
	Field     string `json:"field,omitempty"`
	Value     string `json:"value,omitempty"`
	Condition int    `json:"-"`
}
