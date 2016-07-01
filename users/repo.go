package users

import (

	"fmt"
	"os"
	"database/sql"
	"log"

	"github.com/eirwin/polling-machine/models"

	_ "github.com/lib/pq"
)

type Repo interface  {
	Create(email,password string) (models.User,error)
	Get(id string) (models.User,error)

}

type userRepository struct  {
}

func (u *userRepository) Create(id,email,password string) (models.User,error)  {
	_, err := getDatabase()
	if err != nil {
		log.Fatal(err)
	}

	//var userid int
	//err := db.QueryRow(`INSERT INTO users(name, favorite_fruit, age)
	//VALUES('beatrice', 'starfruit', 93) RETURNING id`).Scan(&userid)

	var user models.User

	//_, err = db.Exec("insert into mydata(val) values(0)")
	if err != nil {
		log.Fatal(err)
	}

	return  user
}

func (u *userRepository) Get(id string)(models.User,error)  {
	_, err := getDatabase()
	if err != nil {
		log.Fatal(err)
	}

	var user models.User

	//rows, err := db.Query("SELECT * FROM users WHERE id = $1",id)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for rows.Next() {
	//	var id int
	//
	//	err = rows.Scan(&id)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	//fmt.Fprintf(resp, "ID: %d\n", id)
	//}

	return  user
}

func getDatabase() (*sql.DB,error){
	connInfo := fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		"postgres",
		"postgres",
		os.Getenv("DB_ENV_POSTGRES_PASSWORD"),
		os.Getenv("POLL-MACHINE_POSTGRES_1_PORT_5432_TCP_ADDR"),
		os.Getenv("POLL-MACHINE_POSTGRES_1_PORT_5432_TCP_PORT"),
	)

	log.Println(connInfo)

	var err error
	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		log.Fatal(err)
	}

	return  db
}

func NewRepo() (Repo,error)   {
	return  &userRepository{},nil
}