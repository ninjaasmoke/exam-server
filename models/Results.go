package models

import "gorm.io/gorm"

// results of a user for each test

type Result struct {
	gorm.Model
	ID      uint   `json:"id" gorm:"primary_key"`
	UserID  uint   `json:"user_id"`
	TestID  uint   `json:"test_id"`
	Marks   uint   `json:"marks"`
	Result  string `json:"result"`
	Deleted bool   `json:"deleted"`
}
