package polls

import (
	"time"
	"github.com/jinzhu/gorm"
)

type createPollRequest struct  {
	Title  string    `json:"title"`
	End    time.Time `json:"end"`
	UserID int       `json:"user_id"`
}

type createPollResponse struct  {
	gorm.Model
	Title  string    `json:"title"`
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
	UserID int       `json:"user_id"`
}