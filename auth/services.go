package auth

import (

)
import (
	"github.com/eirwin/polling-machine/users"
	"log"
	"github.com/dgrijalva/jwt-go"
	"time"
)
var mySigningKey = []byte("GQDstcKsx0NHjPOuXOYg5MbeJ1XT0uFiwDVvVBrk")
//
//const (
//	JWT_SECRET = []byte("GQDstcKsx0NHjPOuXOYg5MbeJ1XT0uFiwDVvVBrk")
//)

type Service interface  {
	Login(email,password string) (string,error)
}

type service struct  {
	users users.Repo
}

func (s *service) Login(email,password string) (string,error)  {

	params := make(map[string]interface{})
	params["email"] = email
	params["password"] = password
	users,err := s.users.Find(params)

	if err != nil {

		log.Println(err)
		return  "",err
	}

	if len(users) == 0 || len(users) > 1 {
		return "",err
	}

	user := users[0]

	//generate JWT Token Claims and Sign
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"user_id": user.ID,
		"exp":time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		log.Println(err)
		return "", err
	}

	return tokenString,nil
}

func NewService() Service  {
	users,err := users.NewRepo()
	if err != nil {
		log.Fatal(err)
	}
	return &service{
		users:users,
	}
}