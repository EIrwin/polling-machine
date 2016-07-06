package main

import (
	"github.com/codegangsta/negroni"
	"github.com/eirwin/polling-machine/data"
	"github.com/eirwin/polling-machine/polls"
	"github.com/eirwin/polling-machine/users"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/thoas/stats"

	"github.com/eirwin/polling-machine/auth"
	"github.com/eirwin/polling-machine/models"
	"log"
)

func main() {

	//initialize database if needed
	initDatabase()

	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)

	statsMiddleware := stats.New()

	// create the router
	router := mux.NewRouter()

	// router handler for login
	router.HandleFunc(auth.LoginPath, auth.LoginHandler).Methods("POST")

	// route handler for retrieving users by id
	router.HandleFunc(users.UsersByID, users.GetUserByIdHandler).Methods("GET")

	// route handler for creating users
	router.HandleFunc(users.UserPath, users.CreateUserHandler).Methods("POST")

	// route handler for retrieving polls by id
	router.HandleFunc(polls.PollsByID, polls.GetPollByIDHandler).Methods("GET")

	//router handler for updating polls
	router.HandleFunc(polls.PollsByID, polls.UpdatePollHandler).Methods("PUT")

	// route handler for retrieving polls by user id
	router.HandleFunc(polls.PollsPath, polls.GetPollsByUserIDHandler).Methods("GET")

	// route handler for creating polls
	router.HandleFunc(polls.PollsPath, polls.CreatePollHandler).Methods("POST")

	// router handler for retrieving poll item by id
	router.HandleFunc(polls.PollItemById, polls.GetPollItemByIDHandler).Methods("GET")

	// router handler for retrieving polls items by poll id
	router.HandleFunc(polls.PollItems, polls.GetPollItemsByPollIDHandler).Methods("GET")

	// router handler for creating poll items
	router.HandleFunc(polls.PollItems, polls.CreatePollItemHandler).Methods("POST")

	// router handler for updating poll items
	router.HandleFunc(polls.PollItemById, polls.UpdatePollItemHandler).Methods("PUT")

	// router handler for delete poll items
	router.HandleFunc(polls.PollItemById, polls.DeletePollItemHandler).Methods("DELETE")

	// router handler for creating poll responses
	router.HandleFunc(polls.PollResponse, polls.CreatePollResponseHandler).Methods("POST")

	// router handler for response counts
	router.HandleFunc(polls.ResponseCount, polls.GetResponseCountsHandler).Methods("GET")

	// router handler for generating response tokens
	router.HandleFunc(polls.ResponseToken, polls.GetResponseTokenHandler).Methods("GET")

	// router handler for initializing database
	router.HandleFunc(data.InitDataPath, data.InitializeDatabaseHandler).Methods("GET")

	// router handler for initialize database health check
	router.HandleFunc(data.HealthCheckPath, data.InitializeDatabaseHealthCheckHandler).Methods("GET")

	// router handler for database connection info
	router.HandleFunc(data.ConnectionInfoPath, data.InitializeDiscoverConnectionHandler).Methods("GET")

	//n.Use(statsMiddleware)
	n.Use(cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))
	n.UseHandler(router)
	n.Use(statsMiddleware)
	n.Run(":8181")

}

//the following is just for testing purposes
//to check if we have all the tables we need
func initDatabase() {

	conn := data.GetConnectionInfo()
	db, _ := data.GetDatabase(conn)
	defer db.Close()

	init := false
	init = !(db.HasTable(models.User{}) &&
		db.HasTable(models.Poll{}) &&
		db.HasTable(models.Item{}) &&
		db.HasTable(models.Response{}))

	if init {
		log.Println("Initializing database...")
		data.InitDB(db)
	}
}
