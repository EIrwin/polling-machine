package models

import "github.com/jinzhu/gorm"

type Item struct {
	gorm.Model
	PollID  int `json:"poll_id,string,omitempty"`
	Value   string `json:"value"`
	Display string `json:"display"`
}
