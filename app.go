package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	//"github.com/thoas/stats"
	"github.com/eirwin/polling-machine/data"
	"github.com/eirwin/polling-machine/polls"
	"github.com/eirwin/polling-machine/users"

	"github.com/eirwin/polling-machine/auth"
)

func main() {

	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)

	//statsMiddleware := stats.New()

	// create the router
	router := mux.NewRouter()

	// router handler for login
	router.HandleFunc(auth.LoginPath,auth.LoginHandler).Methods("POST")

	// route handler for retrieving users by id
	router.HandleFunc(users.UsersByID, users.GetUserByIdHandler).Methods("GET")

	// route handler for creating users
	router.HandleFunc(users.UserPath, users.CreateUserHandler).Methods("POST")

	// route handler for retrieving polls by id
	router.HandleFunc(polls.PollsByID, polls.GetPollByIDHandler).Methods("GET")

	// route handler for retrieving polls by user id
	router.HandleFunc(polls.PollsPath,polls.GetPollsByUserIDHandler).Methods("GET")

	// route handler for creating polls
	router.HandleFunc(polls.PollsPath, polls.CreatePollHandler).Methods("POST")

	// router handler for retrieving poll item by id
	router.HandleFunc(polls.PollItemById, polls.GetPollItemByIDHandler).Methods("GET")

	// router handler for retrieving polls items by poll id
	router.HandleFunc(polls.PollItems,polls.GetPollItemsByPollIDHandler).Methods("GET")

	// router handler for creating poll items
	router.HandleFunc(polls.PollItems, polls.CreatePollItemHandler).Methods("POST")

	// router handler for updating poll items
	router.HandleFunc(polls.PollItemById,polls.UpdatePollItemHandler).Methods("PUT")

	// router handler for delete poll items
	router.HandleFunc(polls.PollItemById,polls.DeletePollItemHandler).Methods("DELETE")

	// router handler for creating poll responses
	router.HandleFunc(polls.PollResponse, polls.CreatePollResponseHandler).Methods("POST")

	// router handler for initializing database
	router.HandleFunc(data.InitDataPath, data.InitializeDatabaseHandler).Methods("GET")

	// router handler for initialize database health check
	router.HandleFunc(data.HealthCheckPath, data.InitializeDatabaseHealthCheckHandler).Methods("GET")

	// router handler for database connection info
	router.HandleFunc(data.ConnectionInfoPath, data.InitializeDiscoverConnectionHandler).Methods("GET")

	//n.Use(statsMiddleware)
	n.Use(cors.New(cors.Options{
		AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
	}))
	n.UseHandler(router)

	n.Run(":8181")

}
