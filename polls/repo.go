package polls

import (
	"database/sql"
	"fmt"
	"os"
	"log"
	"github.com/eirwin/polling-machine/models"
	"time"
)

type Repo interface  {
	//polls
	CreatePoll(id,created_by string,start,end time.Time) (models.Poll,error)
	GetPoll(id string) (models.Poll,error)

	//poll items
	CreateItem(id,poll_id,value,display string) (models.Item,error)
	GetPollItem(id string) (models.Item,error)

	//poll responses
	CreateResponse(id,item_id,ip_address string,timestamp time.Time,) (models.Response,error)
}

type pollRepo struct  {
	
}

func CreatePoll(id,created_by string,start,end time.Time) (models.Poll,error) {
	_,err := getDatabase()
	if err != nil {
		log.Fatal(err)
	}

	var poll models.Poll

	//sql

	return  poll,nil
}

func GetPoll(id string) (models.Poll,error)  {
	_,err := getDatabase()
	if err != nil {
		log.Fatal(err)
	}

	var poll models.Poll

	//sql

	return  poll,nil
}

func CreateItem(id,poll_id,value,display string) (models.Item,error)  {
	_,err := getDatabase()
	if err != nil {
		log.Fatal(err)
	}

	var item models.Item

	//sql

	return  item,nil
}

func GetPollItem(id string) (models.Item,error)  {
	_,err := getDatabase()
	if err != nil {
		log.Fatal(err)
	}

	var item models.Item

	//sql

	return  item,nil
}

func CreateResponse(id,item_id,ip_address string,timestamp time.Time,) (models.Response,error)  {
	_,err := getDatabase()
	if err != nil {
		log.Fatal(err)
	}

	var response models.Response

	//sql

	return  response,nil
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

func NewRepo() (Repo,error) {
	return &pollRepo{},nil
}