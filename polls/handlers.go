package polls

import (
	"encoding/json"
	"github.com/eirwin/polling-machine/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"log"
)

const (
	ByID         = "/{id}"
	ByPollItemID = "/{item_id}"

	// APIBase is the base path for API access
	APIBase = "/api/v1/"

	PollsPath    = APIBase + "polls"
	PollsByID    = APIBase + "polls" + ByID
	PollItems    = APIBase + "polls" + ByID + "/items"
	PollItemById = APIBase + "polls" + ByID + "/items" + ByPollItemID
	PollResponse = APIBase + "polls" + ByID + "/responses"
	ResponseCount = APIBase + "polls" + ByID + "/count"
	ResponseToken = APIBase + "polls" + ByID + "/token"
)

func CreatePollHandler(w http.ResponseWriter, r *http.Request) {
	var poll models.Poll
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&poll); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	service := NewService()

	poll, err := service.CreatePoll(poll.UserID,poll.End,poll.Title)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(poll); err != nil {
		panic(err)
	}
}

func GetPollByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id,err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal("invalid poll id format")
		w.WriteHeader(http.StatusInternalServerError)
	}

	service := NewService()

	poll, err := service.GetPoll(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(poll); err != nil {
		panic(err)
	}
}

func GetPollsByUserIDHandler(w http.ResponseWriter,r *http.Request)  {
	user_id,err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		log.Fatal(err)
	}

	service := NewService()

	polls,err := service.GetPollByUser(user_id)
	if err != nil {
		log.Println(err)
		//w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(polls); err != nil {
		panic(err)
	}
}

func UpdatePollHandler(w http.ResponseWriter,r *http.Request)  {
	var poll models.Poll
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&poll); err != nil {
		log.Println(err)
	}

	service := NewService()

	poll, err := service.UpdatePoll(int(poll.ID),poll.UserID,poll.Start,poll.End,poll.Title)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(poll); err != nil {
		panic(err)
	}
}


func CreatePollItemHandler(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&item); err != nil {
		log.Fatal(err)
	}

	service := NewService()

	item, err := service.CreateItem( item.PollID, item.Value, item.Display)
	if err != nil {
		log.Fatal(err);
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(item); err != nil {
		panic(err)
	}
}

func GetPollItemByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id,err := strconv.Atoi(vars["item_id"])
	if err != nil {
		log.Fatal("invalid poll id format")
		w.WriteHeader(http.StatusInternalServerError)
	}

	service := NewService()

	item, err := service.GetPollItem(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(item); err != nil {
		panic(err)
	}
}

func GetPollItemsByPollIDHandler(w http.ResponseWriter, r * http.Request)  {
	vars := mux.Vars(r)
	id,err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal("invalid poll id format")
		w.WriteHeader(http.StatusInternalServerError)
	}

	service := NewService()

	items, err := service.GetPollItemsByPollID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(items); err != nil {
		panic(err)
	}
}

func UpdatePollItemHandler(w http.ResponseWriter,r * http.Request)  {
	vars := mux.Vars(r)
	id,err := strconv.Atoi(vars["item_id"])

	if err != nil {
		log.Fatal("invalid item id format")
		w.WriteHeader(http.StatusInternalServerError)
	}

	poll_id,err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal("invalid poll id format")
		w.WriteHeader(http.StatusInternalServerError)
	}

	var item models.Item
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&item); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	service := NewService()

	items, err := service.UpdatePollItem(id,poll_id,item.Value,item.Display)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(items); err != nil {
		panic(err)
	}
}

func DeletePollItemHandler(w http.ResponseWriter, r * http.Request)  {
	vars := mux.Vars(r)
	id,err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal("invalid poll id format")
		w.WriteHeader(http.StatusInternalServerError)
	}

	service := NewService()

	items, err := service.GetPollItemsByPollID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(items); err != nil {
		panic(err)
	}
}

func CreatePollResponseHandler(w http.ResponseWriter, r *http.Request) {
	var response models.Response
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&response); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	service := NewService()

	log.Println(response);
	response, err := service.CreateResponse(response.ItemID,response.PollID,r.RemoteAddr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func GetResponseCountsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id,err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal("invalid poll id format")
		w.WriteHeader(http.StatusInternalServerError)
	}

	service := NewService()

	counts, err := service.GetResponseCounts(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(counts); err != nil {
		panic(err)
	}
}

func GetResponseTokenHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id,err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal("invalid poll id format")
		w.WriteHeader(http.StatusInternalServerError)
	}

	service := NewService()

	counts, err := service.GetResponseToken(id)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(counts); err != nil {
		panic(err)
	}
}

