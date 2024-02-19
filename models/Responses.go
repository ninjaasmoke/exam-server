package models

import "gorm.io/gorm"

// all responses for all questions in a test by each user

type Response struct {
	gorm.Model
	UserID     uint   `json:"user_id" gorm:"non null;uniqueIndex:unique_response;foreignKey:UserID"`
	TestID     uint   `json:"test_id" gorm:"non null;uniqueIndex:unique_response;foreignKey:TestID"`
	QuestionID uint   `json:"question_id" gorm:"non null;uniqueIndex:unique_response;foreignKey:QuestionID"`
	Response   string `json:"response"`
	Deleted    bool   `json:"deleted"`
}
