package models

import "time"

type Response struct {
	ID         string    `json:"id",validate:len=32`
	PollItemID string    `json:"poll_item_id"`
	Timestamp  time.Time `json:"timestamp"`
	IpAddress  string    `json:ip_address"`
}
