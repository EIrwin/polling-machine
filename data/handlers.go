package data

import (
	"encoding/json"
	"github.com/eirwin/polling-machine/models"
	"github.com/jinzhu/gorm"
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

func InitializeDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	conn := GetConnectionInfo()
	db, _ := GetDatabase(conn)
	defer db.Close()

	InitDB(db)
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

func InitDB(db *gorm.DB) {
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
