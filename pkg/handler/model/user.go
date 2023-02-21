package model

type UserRegistrationInput struct {
	Name           string `json:"name" binding:"required,min=3,max=64"`
	PhoneNumber    string `json:"phoneNumber" binding:"required,min=6,max=9"`
	Email          string `json:"email" binding:"required,email,max=32"`
	Password       string `json:"password" binding:"required"`
	PasswordRepeat string `json:"passwordRepeat" binding:"required"`
}

type UserLoginInput struct {
	Username    string `json:"username,omitempty" binding:"max=64"`
	PhoneNumber string `json:"phoneNumber,omitempty" binding:"max=64"`
	Password    string `json:"password" binding:"required,min=3,max=64"`
}
