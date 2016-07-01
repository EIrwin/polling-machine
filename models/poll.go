package models

import "time"

type Poll struct {
	ID        string    `json:"id",validate:len=32`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
	CreatedBy string    `json:"created_by"`
}
