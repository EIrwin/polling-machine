package models

import "github.com/jinzhu/gorm"

type Item struct {
	gorm.Model
	PollID  string `json:"poll_id"`
	Value   string `json:"value"`
	Display string `json:"display"`
}
