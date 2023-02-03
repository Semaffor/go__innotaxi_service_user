package model

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber int    `json:"phoneNumber"`
	Email       int    `json:"email"`
	Password    int    `json:"password"`
}
