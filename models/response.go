package models

import (
	"github.com/jinzhu/gorm"
)

type Response struct {
	gorm.Model
	ItemID    int `json:"item_id"`
	PollID	  int `json:"poll_id,string"`
	Token	  string `json:"token" gorm:"-"`
}

type ResponseCount struct {
	ItemID	   int 	  `json:"item_id"`
	Display    string `json:"display"`
	Value	   string `json:"value""`
	Count      int    `json:"count"`
}
