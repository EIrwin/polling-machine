package polls

import (
	"github.com/eirwin/polling-machine/models"
	"log"
	"time"
)

type Service interface {
	//polls
	CreatePoll(id, created_by string, start, end time.Time) (models.Poll, error)
	GetPoll(id string) (models.Poll, error)

	//poll items
	CreateItem(id, poll_id, value, display string) (models.Item, error)
	GetPollItem(id string) (models.Item, error)

	//poll responses
	CreateResponse(id, item_id, ip_address string, timestamp time.Time) (models.Response, error)
}

type service struct {
	polls Repo
}

func (s *service) CreatePoll(id, created_by string, start, end time.Time) (models.Poll, error) {
	poll, err := s.polls.CreatePoll(id, created_by, start, end)
	if err != nil {
		log.Fatal(err)
	}
	return poll, nil
}

func (s *service) GetPoll(id string) (models.Poll, error) {
	poll, err := s.polls.GetPoll(id)
	if err != nil {
		log.Fatal(err)
	}
	return poll, nil
}

func (s *service) CreateItem(id, poll_id, value, display string) (models.Item, error) {
	item, err := s.polls.CreateItem(id, poll_id, value, display)
	if err != nil {
		log.Fatal(err)
	}
	return item, nil
}

func (s *service) GetPollItem(id string) (models.Item, error) {
	item, err := s.polls.GetPollItem(id)
	if err != nil {
		log.Fatal(err)
	}
	return item, nil
}

func (s *service) CreateResponse(id, item_id, ip_address string, timestamp time.Time) (models.Response, error) {
	response, err := s.polls.CreateResponse(id, item_id, ip_address, timestamp)
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
