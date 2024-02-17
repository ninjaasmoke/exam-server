package models

import "gorm.io/gorm"

// a class of users where each user is a `student`

type Class struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Students []User `json:"students" gorm:"many2many:class_students;"`
}
