package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	UserName  string     `json:"username" gorm:"column:username"`
	Role      uint       `json:"role"`
	Deleted   bool       `json:"deleted"`
	Responses []Response `gorm:"foreignKey:UserID"`
}
