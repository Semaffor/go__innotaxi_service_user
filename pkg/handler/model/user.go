package model

type UserRegistrationInput struct {
	Name           string `json:"name" binding:"required,min=3,max=64"`
	PhoneNumber    string `json:"phoneNumber" binding:"required,min=6,max=9"`
	Email          string `json:"email" binding:"required,email,max=32"`
	Password       string `json:"password" binding:"required"`
	PasswordRepeat string `json:"passwordRepeat" binding:"required"`
}

type UserUpdateInput struct {
	Id             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty" binding:"max=64"`
	Username       string `json:"username,omitempty" binding:"max=64"`
	PhoneNumber    string `json:"phoneNumber,omitempty" binding:"max=9"`
	Email          string `json:"email,omitempty" binding:"email,max=32"`
	Password       string `json:"password,omitempty" binding:"max=32"`
	PasswordRepeat string `json:"passwordRepeat,omitempty" binding:"max=32"`
}

type UserLoginInput struct {
	Username    string `json:"username,omitempty" binding:"max=64"`
	PhoneNumber string `json:"phoneNumber,omitempty" binding:"max=64"`
	Password    string `json:"password" binding:"required,min=3,max=64"`
}
