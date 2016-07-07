package users

import (
	"github.com/eirwin/code-challenge/models"
	"github.com/jinzhu/gorm"
)

type createUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type createUserResponse struct {
	gorm.Model
	Error string `json:"error"`
	Email string `json:"email"`
}

type getUserResponse struct {
	Error string `json:"error"`
	models.User
}

func (r *createUserRequest) Validate() (bool, string) {
	var msg string
	if len(r.Email) == 0 {
		return false, "Please enter email"
	}

	if len(r.Password) == 0 {
		return false, "Please enter password"
	}

	return true, msg
}