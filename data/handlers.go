package data

import (
	"net/http"
	"github.com/jinzhu/gorm"
	"fmt"
	"log"
	"os"
	"github.com/eirwin/polling-machine/models"
	"encoding/json"
)

const (

	HealthCheck = "/ping"
	Connection = "/info"

	// APIBase is the base path for API access
	APIBase = "/api/v1/"

	DataPath = APIBase + "data"

	HealthCheckPath = APIBase + "data" + HealthCheck

	ConnectionInfoPath = APIBase + "data" + Connection
)

type ConnectionInfo struct {
	User string
	DB     string
	Password string
	Host string
	Port string
}

func InitializeDatabaseHandler(w http.ResponseWriter,r *http.Request)  {
	conn := getConnectionInfo()
	db := getDatabase(conn)
	initDB(db)
}

func InitializeDatabaseHealthCheckHandler(w http.ResponseWriter,r *http.Request)  {
	conn := getConnectionInfo()
	_ = getDatabase(conn)
	w.WriteHeader(http.StatusOK)
}

func InitializeDiscoverConnectionHandler(w http.ResponseWriter,r *http.Request)  {
	conn := getConnectionInfo()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(conn); err != nil {
		panic(err)
	}
}

func initDB(db *gorm.DB)  {
	createUsersTable(db)
	createPollsTable(db)
	createItemsTable(db)
	createResponseTable(db)
}

func createUsersTable(db *gorm.DB)  {
	db.DropTableIfExists(&models.User{})
	db.CreateTable(&models.User{})
}

func createPollsTable(db *gorm.DB)  {
	db.DropTableIfExists(&models.Poll{})
	db.CreateTable(&models.Poll{})
}

func createItemsTable(db *gorm.DB){
	db.DropTableIfExists(&models.Item{})
	db.CreateTable(&models.Item{})
}

func createResponseTable(db *gorm.DB)  {
	db.DropTableIfExists(&models.Response{})
	db.CreateTable(&models.Response{})
}

func getDatabase(connInfo ConnectionInfo) (*gorm.DB)  {
	conn := fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		connInfo.User,
		connInfo.DB,
		connInfo.Password,
		connInfo.Host,
		connInfo.Port,
	)

	db,err := gorm.Open("postgres",conn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func getConnectionInfo() ConnectionInfo {
	return ConnectionInfo{
		User : "postgres",
		DB : "postgres",
		Password: "mypass",
		Host:os.Getenv("POLLINGMACHINE_POSTGRES_1_PORT_5432_TCP_ADDR"),
		Port:os.Getenv("POLLINGMACHINE_1_PORT_5432_TCP_PORT"),
	}
}