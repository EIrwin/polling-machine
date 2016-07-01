package users

import (
	"github.com/eirwin/polling-machine/models"

	"log"
)

type Service interface  {
	Create(id,email,password string) (models.User,error)
	Get(id string) (models.User,error)
}

type service struct {
	users Repo
}

func (s *service) Create(id,email,password string) (models.User,error)  {
	user,err := s.users.Create(id,email,password)
	if err != nil {
		log.Fatal(err)
	}
	return  user,nil
}

func (s *service) Get(id string) (models.User,error)  {
	user,err := s.Get(id)
	if err != nil {
		log.Fatal(err)
	}
	return  user,nil
}

//NewService returns a new instance of the Service with dependencies
func NewService() Service  {
	users,err := NewRepo()
	if err != nil {
		log.Fatal(err)
	}

	return &service {
		users : users,
	}
}

