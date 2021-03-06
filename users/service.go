package users

import (
	"github.com/eirwin/polling-machine/models"

	"log"
)

type Service interface {
	Create(email, password string) (models.User, error)
	Get(id int) (models.User, error)
	GetByEmail(email string) (models.User, error)
}

type service struct {
	users Repo
}

func (s *service) Create(email, password string) (models.User, error) {
	user, err := s.users.Create(email, password)
	if err != nil {

		if err.Error() == "duplicate email" {
			return user, err
		}

		log.Fatal(err)
	}
	return user, nil
}

func (s *service) Get(id int) (models.User, error) {
	user, err := s.users.Get(id)
	if err != nil {
		log.Fatal(err)
	}
	return user, nil
}

func (s *service) GetByEmail(email string) (models.User, error) {

	//search by email
	params := make(map[string]interface{})
	params["email"] = email

	users, err := s.users.Find(params)
	if err != nil {
		log.Fatal(err)
	}

	return users[0], nil
}

//NewService returns a new instance of the Service with dependencies
func NewService() Service {
	users, err := NewRepo()
	if err != nil {
		log.Fatal(err)
	}

	return &service{
		users: users,
	}
}
