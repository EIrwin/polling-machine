package data

import (
	"encoding/json"
	"fmt"
	"github.com/eirwin/polling-machine/models"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

const (
	InitData    = "/init"
	HealthCheck = "/ping"
	Connection  = "/info"

	// APIBase is the base path for API access
	APIBase = "/api/v1/"

	InitDataPath = APIBase + "data" + InitData

	HealthCheckPath = APIBase + "data" + HealthCheck

	ConnectionInfoPath = APIBase + "data" + Connection
)

type ConnectionInfo struct {
	User     string
	DB       string
	Password string
	Host     string
	Port     string
}

func InitializeDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	conn := GetConnectionInfo()
	db, _ := GetDatabase(conn)
	initDB(db)
}

func InitializeDatabaseHealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	conn := GetConnectionInfo()
	_, _ = GetDatabase(conn)

	w.WriteHeader(http.StatusOK)
}

func InitializeDiscoverConnectionHandler(w http.ResponseWriter, r *http.Request) {
	conn := GetConnectionInfo()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(conn); err != nil {
		panic(err)
	}
}

func initDB(db *gorm.DB) {
	createUsersTable(db)
	createPollsTable(db)
	createItemsTable(db)
	createResponseTable(db)
}

func createUsersTable(db *gorm.DB) {
	db.DropTableIfExists(&models.User{})
	db.CreateTable(&models.User{})
}

func createPollsTable(db *gorm.DB) {
	db.DropTableIfExists(&models.Poll{})
	db.CreateTable(&models.Poll{})
}

func createItemsTable(db *gorm.DB) {
	db.DropTableIfExists(&models.Item{})
	db.CreateTable(&models.Item{})
}

func createResponseTable(db *gorm.DB) {
	db.DropTableIfExists(&models.Response{})
	db.CreateTable(&models.Response{})
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
