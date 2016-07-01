package models

type User struct {
	ID       string    `json:"id",validate:len=32`
	Email 	 string    `json:"email"`
	Password string    `json:"password"`
}
