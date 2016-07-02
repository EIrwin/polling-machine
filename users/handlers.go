package users

import (
	"net/http"
	"encoding/json"
	"log"

	"github.com/eirwin/polling-machine/models"

	"github.com/pborman/uuid"
	"github.com/gorilla/mux"
)

const (
	ByID = "/{id}"

	// APIBase is the base path for API access
	APIBase = "/api/v1/"

	UserPath = APIBase + "users"
	UsersByID = APIBase + "users" + ByID
)

//CreateUserHandler
func CreateUserHandler(w http.ResponseWriter, r *http.Request)  {
	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	service := NewService()
	user,err := service.Create(
		uuid.NewUUID().String(),
		user.Email,
		user.Password)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

//GetUserByIdHandler
func GetUserByIdHandler(w http.ResponseWriter,r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["id"]

	service := NewService()
	user,err := service.Get(id)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}

}
