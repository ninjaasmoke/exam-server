package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Role     uint   `json:"role"`
	UserID   uint   `json:"user_id"`
	UserName string `json:"username" gorm:"column:username"`
	jwt.StandardClaims
}
