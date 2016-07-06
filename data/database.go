package data

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type ConnectionInfo struct {
	User     string
	DB       string
	Password string
	Host     string
	Port     string
}

func GetDatabase(connInfo ConnectionInfo) (*gorm.DB, error) {
	conn := fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		connInfo.User,
		connInfo.DB,
		connInfo.Password,
		connInfo.Host,
		connInfo.Port,
	)

	db, err := gorm.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func GetConnectionInfo() ConnectionInfo {
	return ConnectionInfo{
		User:     "postgres",
		DB:       "postgres",
		Password: "mypass",
		//Host:"172.17.0.2",
		Host: "192.168.99.100",
		Port: "5432",
	}
	//return ConnectionInfo{
	//	User : "postgres",
	//	DB : "postgres",
	//	Password: "mypass",
	//	Host:os.Getenv("POLLINGMACHINE_POSTGRES_1_PORT_5432_TCP_ADDR"),
	//	Port:os.Getenv("POLLINGMACHINE_POSTGRES_1_PORT_5432_TCP_PORT"),
	//}
}
