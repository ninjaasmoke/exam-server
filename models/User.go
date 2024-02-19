package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID    uint       `json:"user_id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	UserName  string     `json:"username" gorm:"column:username"`
	Role      uint       `json:"role"`
	Deleted   bool       `json:"deleted"`
	Responses []Response `gorm:"foreignKey:UserID"`
}
