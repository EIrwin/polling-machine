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
	Find(params map[string]interface{}) ([]models.User,error)
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
	//user := models.User{}
	//user.ID = uint(id)
	//
	//var polls []models.Poll
	//var items []models.Item
	//
	//db.Model(&user).Related(&polls).Related(&items)


	return user, nil
}

func (u *userRepository) Find(params map[string]interface{}) ([]models.User,error)  {
	connInfo := data.GetConnectionInfo()
	db, err := data.GetDatabase(connInfo)
	if err != nil {
		log.Print("here")
		log.Fatal(err)
	}

	var users []models.User
	db.Where(params).Find(&users)

	return users, nil
}

func NewRepo() (Repo, error) {
	return &userRepository{}, nil
}
