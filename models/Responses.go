package models

import "gorm.io/gorm"

// all responses for all questions in a test by each user

type Response struct {
	gorm.Model
	UserID     uint   `json:"user_id"`
	TestID     uint   `json:"test_id"`
	QuestionID uint   `json:"question_id"`
	Response   string `json:"response"`
	Deleted    bool   `json:"deleted"`
}
