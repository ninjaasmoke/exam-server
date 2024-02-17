package models

import (
	"gorm.io/gorm"
)

// a test consists of a set of questions

type Test struct {
	gorm.Model
	Name         string     `json:"name" gorm:"not null"`
	Subject      string     `json:"subject"`
	Chapter      string     `json:"chapter"`
	Level        string     `json:"level"`
	Questions    []Question `json:"questions" gorm:"many2many:test_questions;not null"`
	TotalMarks   int        `json:"total_marks"`
	TotalTime    int        `json:"total_time"`
	Instructions string     `json:"instructions"`
	Deleted      bool       `json:"deleted"`
}
