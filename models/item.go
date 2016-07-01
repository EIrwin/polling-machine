package models

type Item struct  {
	ID       string    `json:"id",validate:len=32`
	PollID	 string	   `json:"poll_id"`
	Value    string	   `json:"value"`
	Display  string	   `json:"display"`
}