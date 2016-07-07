package polls

import (
	"encoding/json"
	"github.com/eirwin/polling-machine/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

const (
	ByID         = "/{id}"
	ByPollItemID = "/{item_id}"

	// APIBase is the base path for API access
	APIBase = "/api/v1/"

	PollsPath     = APIBase + "polls"
	PollsByID     = APIBase + "polls" + ByID
	PollItems     = APIBase + "polls" + ByID + "/items"
	PollItemById  = APIBase + "polls" + ByID + "/items" + ByPollItemID
	PollResponse  = APIBase + "polls" + ByID + "/responses"
	ResponseCount = APIBase + "polls" + ByID + "/count"
	ResponseToken = APIBase + "polls" + ByID + "/token"
)

func CreatePollHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var request createPollRequest
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&request); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if valid,msg := request.Validate(); !valid {
		w.WriteHeader(http.StatusInternalServerError)
		response :=  createPollResponse{
			Error:msg,
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	service := NewService()

	poll, err := service.CreatePoll(request.UserID,request.End,request.Title)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(poll); err != nil {
		log.Fatal(err)
	}
}

func GetPollByIDHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := NewService()

	poll, err := service.GetPoll(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(poll); err != nil {
		panic(err)
	}
}

func GetPollsByUserIDHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	user_id, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		log.Fatal(err)
	}

	service := NewService()

	polls, err := service.GetPollByUser(user_id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(polls); err != nil {
		log.Println(err)
	}
}

func UpdatePollHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var request updatePollRequest
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&request); err != nil {
		log.Println(err)
	}

	if valid,msg := request.Validate(); !valid {
		w.WriteHeader(http.StatusInternalServerError)
		response :=  updatePollResponse{
			Error:msg,
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	service := NewService()

	poll, err := service.UpdatePoll(int(request.ID),request.UserID,request.Start,request.End, request.Title)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(poll); err != nil {
		log.Println(err)
	}
}

func CreatePollItemHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var request createPollItemRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		log.Fatal(err)
	}

	if valid,msg := request.Validate(); !valid {
		w.WriteHeader(http.StatusInternalServerError)
		response :=  createPollItemResponse{
			Error:msg,
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	service := NewService()

	item, err := service.CreateItem(request.PollID,request.Value,request.Display)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(item); err != nil {
		log.Println(err)
	}
}

func GetPollItemByIDHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["item_id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := NewService()

	item, err := service.GetPollItem(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(item); err != nil {
		log.Println(err)
	}
}

func GetPollItemsByPollIDHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := NewService()

	items, err := service.GetPollItemsByPollID(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(items); err != nil {
		log.Println(err)
	}
}

func UpdatePollItemHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["item_id"])

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	poll_id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var item models.Item
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&item); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := NewService()

	items, err := service.UpdatePollItem(id, poll_id, item.Value, item.Display)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(items); err != nil {
		log.Println(err)
	}
}

func DeletePollItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := NewService()

	items, err := service.GetPollItemsByPollID(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(items); err != nil {
		log.Println(err)
	}
}

func CreatePollResponseHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var request createPollResponseRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if valid,msg := request.Validate(); !valid {
		w.WriteHeader(http.StatusInternalServerError)
		response :=  createPollResponseResponse{
			Error:msg,
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	service := NewService()

	response, err := service.CreateResponse(request.ItemID, request.PollID, r.RemoteAddr)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println(err)
	}
}

func GetResponseCountsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := NewService()

	counts, err := service.GetResponseCounts(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(counts); err != nil {
		log.Println(err)
	}
}

func GetResponseTokenHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := NewService()

	counts, err := service.GetResponseToken(id)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(counts); err != nil {
		log.Println(err)
	}
}
