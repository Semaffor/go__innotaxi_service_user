package domain

import "github.com/dgrijalva/jwt-go"

type JwtTokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UserCredentials struct {
	Username     string `json:"username,omitempty"`
	MobileNumber string `json:"mobileNumber,omitempty"`
	Password     string `json:"password,required"`
}

type JwtClaims struct {
	jwt.StandardClaims
	UserId int    `json:"userId,required"`
	Role   string `json:"role,required"`
}
