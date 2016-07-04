package polls

import (
	"github.com/eirwin/polling-machine/models"
	"log"
	"time"
)

type Service interface {
	//polls
	CreatePoll(user_id int,end time.Time,title string) (models.Poll, error)
	GetPoll(id int) (models.Poll, error)
	GetPollByUser(user_id int) ([]models.Poll,error)

	//poll items
	CreateItem(poll_id int, value, display string) (models.Item, error)
	GetPollItem(id int) (models.Item, error)
	GetPollItemsByPollID(poll_id int) ([]models.Item,error)
	UpdatePollItem(id,poll_id int,value,display string) (models.Item,error)
	DeleteItem(id int) (error)


	//poll responses
	CreateResponse(item_id int, ip_address string) (models.Response, error)
}

type service struct {
	polls Repo
}

func (s *service) CreatePoll(user_id int,end time.Time,title string) (models.Poll, error) {
	poll, err := s.polls.CreatePoll(user_id,time.Now(), end,title)
	if err != nil {
		log.Fatal(err)
	}
	return poll, nil
}

func (s *service) GetPoll(id int) (models.Poll, error) {
	poll, err := s.polls.GetPoll(id)
	if err != nil {
		log.Fatal(err)
	}
	return poll, nil
}

func (s *service) GetPollByUser(user_id int) ([]models.Poll,error)  {
	polls,err := s.polls.GetPollsByUser(user_id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(polls)
	return  polls,nil
}

func (s *service) CreateItem(poll_id int, value, display string) (models.Item, error) {
	item, err := s.polls.CreateItem(poll_id, value, display)
	log.Println(item)
	if err != nil {
		log.Fatal(err)
	}
	return item, nil
}

func (s *service) GetPollItem(id int) (models.Item, error) {
	item, err := s.polls.GetPollItem(id)
	if err != nil {
		log.Fatal(err)
	}
	return item, nil
}

func (s *service) GetPollItemsByPollID(poll_id int) ([]models.Item,error)  {
	items,err := s.polls.GetPollItemsByPollID(poll_id)
	if err != nil {
		log.Fatal(err)
	}
	return items,nil
}

func (s *service) UpdatePollItem(id,poll_id int,value,display string) (models.Item,error)  {
	item,err := s.polls.UpdatePollItem(id,poll_id,value,display)
	if err != nil {
		log.Println("here")
		log.Fatal(err)
	}
	return item,nil
}
func (s *service) DeleteItem(id int) (error)  {
	err := s.polls.DeleteItem(id)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
func (s *service) CreateResponse(item_id int, ip_address string) (models.Response, error) {
	response, err := s.polls.CreateResponse(item_id, ip_address)
	if err != nil {
		log.Fatal(err)
	}
	return response, nil
}

func NewService() Service {
	polls, err := NewRepo()
	if err != nil {
		log.Fatal(err)
	}

	return &service{
		polls: polls,
	}
}
