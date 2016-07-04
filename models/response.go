package models

import (
	"github.com/jinzhu/gorm"
)

type Response struct {
	gorm.Model
	ItemID    int `json:"item_id"`
	IpAddress string `json:ip_address"`
}
