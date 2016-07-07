package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"strconv"
)

const (
	ByID = "/{id}"

	// APIBase is the base path for API access
	APIBase = "/api/v1/"

	UserPath  = APIBase + "users"
	UsersByID = APIBase + "users" + ByID
)

//CreateUserHandler
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var request createUserRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if valid, msg := request.Validate(); !valid {
		w.WriteHeader(http.StatusBadRequest)
		response := &createUserResponse{
			Error: msg,
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	service := NewService()
	user, err := service.Create(
		request.Email,
		request.Password)

	if err != nil {
		if err.Error() == "duplicate email" {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

//GetUserByIdHandler
func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := NewService()
	user, err := service.Get(id)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Println(err)
		return
	}
}
