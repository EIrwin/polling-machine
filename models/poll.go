package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Poll struct {
	gorm.Model
	Title  string    `json:"title"`
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
	UserID int       `json:"user_id"`
}
