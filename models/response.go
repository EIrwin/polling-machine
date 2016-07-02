package models

import (
	"github.com/jinzhu/gorm"
)

type Response struct {
	gorm.Model
	ItemID    string `json:"item_id"`
	IpAddress string `json:ip_address"`
}
