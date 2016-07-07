package polls

import (
	"github.com/jinzhu/gorm"
	"time"
)

type createPollRequest struct {
	Title  string    `json:"title"`
	End    time.Time `json:"end"`
	UserID int       `json:"user_id"`
}

type createPollResponse struct {
	gorm.Model
	Error  string    `json"error"`
	Title  string    `json:"title"`
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
	UserID int       `json:"user_id"`
}

type updatePollRequest struct {
	gorm.Model
	Title  string    `json:"title"`
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
	UserID int       `json:"user_id"`
}

type updatePollResponse struct {
	gorm.Model
	Error  string    `json:"error"`
	Title  string    `json:"title"`
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
	UserID int       `json:"user_id"`
}

type createPollItemRequest struct {
	PollID  int    `json:"poll_id,string,omitempty"`
	Value   string `json:"value"`
	Display string `json:"display"`
}

type createPollItemResponse struct {
	gorm.Model
	Error   string `json:"error"`
	PollID  int    `json:"poll_id,string,omitempty"`
	Value   string `json:"value"`
	Display string `json:"display"`
}

type updatePollItemRequest struct {
	PollID  int    `json:"poll_id,string,omitempty"`
	Value   string `json:"value"`
	Display string `json:"display"`
}

type updatePollItemResponse struct {
	gorm.Model
	Error   string `json:"error"`
	PollID  int    `json:"poll_id,string,omitempty"`
	Value   string `json:"value"`
	Display string `json:"display"`
}

type createPollResponseRequest struct  {
	ItemID int    `json:"item_id"`
	PollID int    `json:"poll_id,string"`
	Token  string `json:"token" gorm:"-"`
}

type createPollResponseResponse struct  {
	gorm.Model
	ItemID int    `json:"item_id"`
	PollID int    `json:"poll_id,string"`
}

func (r *createPollRequest) Validate() (bool, string) {
	var msg string
	if len(r.Title) == 0 {
		return false, "Title cannot be empty"
	}

	if r.UserID <= 0 {
		return false, "Invalid UserID"
	}
	return true, msg
}

func (r *updatePollRequest) Validate() (bool, string) {
	var msg string
	if len(r.Title) == 0 {
		return false, "Title cannot be empty"
	}

	if r.UserID <= 0 {
		return false, "Invalid UserID"
	}

	return true, msg
}

func (r *createPollItemRequest) Validate() (bool, string) {
	var msg string
	if len(r.Value) == 0 {
		return false, "Value cannot be empty"
	}

	if len(r.Display) == 0 {
		return false, "Display cannot be empty"
	}

	if r.PollID <= 0 {
		return false, "Invalid PollID"
	}
	return true, msg
}

func (r *updatePollItemRequest) Validate() (bool,string)  {
	var msg string
	if len(r.Value) == 0 {
		return false, "Value cannot be empty"
	}

	if len(r.Display) == 0 {
		return false, "Display cannot be empty"
	}

	if r.PollID <= 0 {
		return false, "Invalid PollID"
	}
	return true, msg
}

func (r *createPollResponseRequest) Validate() (bool,string)  {
	var msg string

	if len(r.Token) == 0 {
		return false, "Token cannot be empty"
	}

	if r.PollID <= 0 {
		return false, "Invalid PollID"
	}

	if r.ItemID <= 0 {
		return  false,"Invalid ItemID"
	}

	return true, msg
}