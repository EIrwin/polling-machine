package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Poll struct {
	gorm.Model
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
	UserID string    `json:"created_by"`
}
