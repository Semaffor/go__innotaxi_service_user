package model

type User struct {
	Id           int    `json:"id"`
	Username     string `json:"name"`
	PhoneNumber  string `json:"phoneNumber"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`
}
