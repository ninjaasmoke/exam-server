package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Role uint `json:"role"`
	jwt.StandardClaims
}
