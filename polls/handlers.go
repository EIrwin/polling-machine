package polls

import (
	"net/http"
	"github.com/eirwin/polling-machine/models"
	"encoding/json"
	"github.com/pborman/uuid"
	"github.com/gorilla/mux"
	"time"
)

const (
	ByID = "/{id}"
	ByPollItemID = "/{item_id}"

	// APIBase is the base path for API access
	APIBase = "/api/v1/"

	PollsPath = APIBase + "polls"
	PollsByID = APIBase + "polls" + ByID
	PollItems = APIBase + "polls" + ByID + "/items"
	PollItemById = APIBase + "polls" + ByID + "/items" + ByPollItemID
	PollResponse = APIBase + "polls" + ByID + "/responses"
)

func CreatePollHandler(w http.ResponseWriter,r *http.Request)  {
	var poll models.Poll
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&poll); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	service := NewService()

	poll,err := service.CreatePoll(uuid.NewUUID().String(),poll.CreatedBy,poll.Start,poll.End)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(poll); err != nil {
		panic(err)
	}
}

func GetPollByIDHandler(w http.ResponseWriter,r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["id"]

	service := NewService()

	poll,err := service.GetPoll(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(poll); err != nil {
		panic(err)
	}
}

func CreatePollItemHandler(w http.ResponseWriter,r *http.Request)  {
	var item models.Item
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&item); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	service := NewService()

	item,err := service.CreateItem(uuid.NewUUID().String(),item.PollID,item.Value,item.Display)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(item); err != nil {
		panic(err)
	}
}

func GetPollItemByIDHandler(w http.ResponseWriter,r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["id"]

	service := NewService()

	item,err := service.GetPollItem(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(item); err != nil {
		panic(err)
	}
}

func CreatePollResponseHandler(w http.ResponseWriter,r *http.Request)  {
	var response models.Response
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	service := NewService()

	response,err := service.CreateResponse(uuid.NewUUID().String(),response.PollItemID,response.IpAddress,time.Now())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
