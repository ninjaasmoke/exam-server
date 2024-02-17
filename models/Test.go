package models

import "gorm.io/gorm"

// a test consists of a set of questions

type Test struct {
	gorm.Model
	ID           uint       `json:"id" gorm:"primary_key"`
	Name         string     `json:"name"`
	Subject      string     `json:"subject"`
	Chapter      string     `json:"chapter"`
	Level        string     `json:"level"`
	Questions    []Question `json:"questions" gorm:"many2many:test_questions;"`
	TotalMarks   int        `json:"total_marks"`
	TotalTime    int        `json:"total_time"`
	Instructions string     `json:"instructions"`
}
