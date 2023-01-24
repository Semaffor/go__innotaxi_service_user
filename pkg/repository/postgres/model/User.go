package model

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber int    `json:"phone_number"`
	Email       int    `json:"email"`
	Password    int    `json:"password"`
}
