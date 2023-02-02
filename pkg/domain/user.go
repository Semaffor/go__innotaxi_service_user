package domain

type User struct {
	Id           int    `json:"id"`
	Username     string `json:"name"`
	PhoneNumber  string `json:"phone_number"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}
