package models

import "gorm.io/gorm"

// A test question which is MCQ and can have anywhere from 2 to 6 options.

type Question struct {
	gorm.Model
	Question string `json:"question"`
	Option1  string `json:"option1"`
	Option2  string `json:"option2"`
	Option3  string `json:"option3"`
	Option4  string `json:"option4"`
	Option5  string `json:"option5"`
	Option6  string `json:"option6"`
	Answer   string `json:"answer"`
	Subject  string `json:"subject"`
	Chapter  string `json:"chapter"`
	Level    string `json:"level"`
	Deleted  bool   `json:"deleted"`
}
