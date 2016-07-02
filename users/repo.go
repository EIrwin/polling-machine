package users

import (
	"log"

	"github.com/eirwin/polling-machine/models"

	"github.com/eirwin/polling-machine/data"
	_ "github.com/lib/pq"
)

type Repo interface {
	Create(email, password string) (models.User, error)
	Get(id int) (models.User, error)
}

type userRepository struct {
}

func (u *userRepository) Create(email, password string) (models.User, error) {
	connInfo := data.GetConnectionInfo()
	db, err := data.GetDatabase(connInfo)
	if err != nil {
		log.Fatal(err)
	}

	user := models.User{
		Email:    email,
		Password: password,
	}

	db.NewRecord(user)
	db.Create(&user)

	if err != nil {
		log.Fatal(err)
	}

	return user, nil
}

func (u *userRepository) Get(id int) (models.User, error) {
	connInfo := data.GetConnectionInfo()
	db, err := data.GetDatabase(connInfo)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	db.First(&user, id)

	return user, nil
}

func NewRepo() (Repo, error) {
	return &userRepository{}, nil
}
