package model

type UserCredentials struct {
	Username     string `json:"username,omitempty"`
	MobileNumber string `json:"mobileNumber,omitempty"`
	Password     string `json:"password"`
}
