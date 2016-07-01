package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/thoas/stats"
	"github.com/eirwin/polling-machine/users"
	"github.com/eirwin/polling-machine/polls"
)

func main(){

	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)

	statsMiddleware := stats.New()

	// create the router
	router := mux.NewRouter()

	// route handler for retrieving users by id
	router.HandleFunc(users.UsersByID,users.GetUserByIdHandler).Methods("GET")

	// route handler for creating users
	router.HandleFunc(users.UserPath,users.CreateUserHandler).Methods("POST")

	// route handler for retrieving polls by id
	router.HandleFunc(polls.PollsByID,polls.GetPollByIDHandler).Methods("GET")

	// route handler for creating polls
	router.HandleFunc(polls.PollsPath,polls.CreatePollHandler).Methods("POST")

	// router handler for retrieving poll item by id
	router.HandleFunc(polls.PollItemById,polls.GetPollItemByIDHandler).Methods("GET")

	// router handler for creating poll items
	router.HandleFunc(polls.PollItems,polls.CreatePollItemHandler).Methods("POST")

	// router handler for creating poll responses
	router.HandleFunc(polls.PollResponse,polls.CreatePollResponseHandler).Methods("POST")

	n.Use(statsMiddleware)
	n.UseHandler(router)
	n.Run(":8181")
}
