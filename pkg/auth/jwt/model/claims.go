package model

import "github.com/dgrijalva/jwt-go"

type JwtClaims struct {
	jwt.StandardClaims
	UserId int    `json:"userId"`
	Role   string `json:"role"`
}
