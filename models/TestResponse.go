package models

import "gorm.io/gorm"

type TestResponse struct {
	gorm.Model
	Name         string     `json:"name"`
	Subject      string     `json:"subject"`
	Chapter      string     `json:"chapter"`
	Level        string     `json:"level"`
	Questions    []Question `json:"questions"`
	TotalMarks   int        `json:"total_marks"`
	TotalTime    int        `json:"total_time"`
	Instructions string     `json:"instructions"`
	Deleted      bool       `json:"deleted"`
}
