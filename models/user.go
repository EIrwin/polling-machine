package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID       string    `json:"id",validate:len=32`
	Email 	 string    `json:"email"`
	Password string    `json:"password"`
	Num          int   `gorm:"AUTO_INCREMENT"`
}
