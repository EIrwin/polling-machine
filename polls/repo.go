package polls

import (
	"github.com/eirwin/polling-machine/data"
	"github.com/eirwin/polling-machine/models"
	"log"
	"time"
)

type Repo interface {
	//polls
	CreatePoll(user_id int, start, end time.Time, title string) (models.Poll, error)
	GetPoll(id int) (models.Poll, error)
	GetPollsByUser(user_id int) ([]models.Poll, error)
	UpdatePoll(id, user_id int, start, end time.Time, title string) (models.Poll, error)

	//poll itemsr
	CreateItem(poll_id int, value, display string) (models.Item, error)
	GetPollItem(id int) (models.Item, error)
	GetPollItemsByPollID(poll_id int) ([]models.Item, error)
	UpdatePollItem(id, poll_id int, value, display string) (models.Item, error)
	DeleteItem(id int) error

	//poll responses
	CreateResponse(item_id, poll_id int) (models.Response, error)
	GetResponseCounts(poll_id int) ([]models.ResponseCount, error)
}

type pollRepo struct {
}

func (repo *pollRepo) CreatePoll(user_id int, start, end time.Time, title string) (models.Poll, error) {
	conn := data.GetConnectionInfo()
	db, err := data.GetDatabase(conn)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	poll := models.Poll{
		Start:  start,
		End:    end,
		UserID: user_id,
		Title:  title,
	}

	db.NewRecord(poll)

	db.Create(&poll)

	return poll, nil
}

func (repo *pollRepo) GetPoll(id int) (models.Poll, error) {
	conn := data.GetConnectionInfo()
	db, err := data.GetDatabase(conn)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var poll models.Poll

	db.First(&poll, id)

	return poll, nil
}

func (repo *pollRepo) GetPollsByUser(user_id int) ([]models.Poll, error) {
	conn := data.GetConnectionInfo()
	db, err := data.GetDatabase(conn)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var polls []models.Poll

	db.Where("user_id = ?", user_id).Find(&polls)
	return polls, nil
}

func (repo *pollRepo) UpdatePoll(id, user_id int, start, end time.Time, title string) (models.Poll, error) {
	conn := data.GetConnectionInfo()
	db, err := data.GetDatabase(conn)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var poll models.Poll
	poll.ID = uint(id)
	poll.UserID = user_id
	poll.Start = start
	poll.End = end
	poll.Title = title

	db.Save(&poll)

	return poll, nil
}

func (repo *pollRepo) CreateItem(poll_id int, value, display string) (models.Item, error) {
	conn := data.GetConnectionInfo()
	db, err := data.GetDatabase(conn)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	item := models.Item{
		PollID:  poll_id,
		Value:   value,
		Display: display,
	}

	db.NewRecord(item)

	db.Create(&item)

	return item, nil
}

func (repo *pollRepo) GetPollItem(id int) (models.Item, error) {
	conn := data.GetConnectionInfo()
	db, err := data.GetDatabase(conn)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var item models.Item
	item.ID = uint(id)
	db.First(&item, id)

	return item, nil
}

func (repo *pollRepo) GetPollItemsByPollID(poll_id int) ([]models.Item, error) {
	conn := data.GetConnectionInfo()
	db, err := data.GetDatabase(conn)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var items []models.Item

	db.Where("poll_id = ?", poll_id).Find(&items)

	return items, nil
}

func (repo *pollRepo) UpdatePollItem(id, poll_id int, value, display string) (models.Item, error) {
	conn := data.GetConnectionInfo()
	db, err := data.GetDatabase(conn)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var item models.Item
	item.ID = uint(id)
	item.PollID = poll_id
	item.Value = value
	item.Display = display

	db.Save(&item)

	return item, nil
}

func (repo *pollRepo) DeleteItem(id int) error {
	conn := data.GetConnectionInfo()
	db, err := data.GetDatabase(conn)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	user.ID = uint(id)
	db.Delete(&user)

	return nil
}

func (repo *pollRepo) CreateResponse(item_id, poll_id int) (models.Response, error) {
	conn := data.GetConnectionInfo()
	db, err := data.GetDatabase(conn)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	response := models.Response{
		ItemID: item_id,
		PollID: poll_id,
	}

	db.NewRecord(response)

	db.Create(&response)

	return response, nil
}

func (repo *pollRepo) GetResponseCounts(poll_id int) ([]models.ResponseCount, error) {
	conn := data.GetConnectionInfo()
	db, err := data.GetDatabase(conn)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	//items
	var items []models.Item
	db.Where("poll_id = ?", poll_id).Find(&items)

	lookup := make(map[int]models.Item)

	//map lookup items
	for key, val := range items {
		lookup[key] = val
	}

	//responses
	var responses []models.Response
	db.Where("poll_id = ?", poll_id).Find(&responses)

	counts := make(map[int]int, len(items))

	for _, val := range responses {
		counts[val.ItemID]++
	}

	responseCounts := make([]models.ResponseCount, len(items))

	for key, val := range lookup {
		response := val
		responseCounts[key] = models.ResponseCount{
			ItemID:  int(response.ID),
			Display: response.Display,
			Value:   response.Value,
			Count:   counts[key],
		}
	}

	return responseCounts, nil
}

func NewRepo() (Repo, error) {
	return &pollRepo{}, nil
}
