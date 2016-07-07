package polls

import (
	"github.com/eirwin/polling-machine/email"
	"github.com/eirwin/polling-machine/models"
	"github.com/eirwin/polling-machine/users"

	"fmt"
	"github.com/eirwin/polling-machine/cache"
	"github.com/pborman/uuid"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Service interface {
	//polls
	CreatePoll(user_id int, end time.Time, title string) (models.Poll, error)
	GetPoll(id int) (models.Poll, error)
	GetPollByUser(user_id int) ([]models.Poll, error)
	UpdatePoll(id, user_id int, start, end time.Time, title string) (models.Poll, error)

	//poll items
	CreateItem(poll_id int, value, display string) (models.Item, error)
	GetPollItem(id int) (models.Item, error)
	GetPollItemsByPollID(poll_id int) ([]models.Item, error)
	UpdatePollItem(id, poll_id int, value, display string) (models.Item, error)
	DeleteItem(id int) error

	//poll responses
	CreateResponse(item_id, poll_id int, token string) (models.Response, error)
	GetResponseCounts(poll_id int) ([]models.ResponseCount, error)

	//keys
	GetResponseToken(pollId int) (string, error)
}

type service struct {
	polls Repo
	users users.Repo
}

func (s *service) CreatePoll(user_id int, end time.Time, title string) (models.Poll, error) {
	poll, err := s.polls.CreatePoll(user_id, time.Now(), end, title)
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

func (s *service) GetPollByUser(user_id int) ([]models.Poll, error) {
	polls, err := s.polls.GetPollsByUser(user_id)
	if err != nil {
		log.Fatal(err)
	}
	return polls, nil
}

func (s *service) UpdatePoll(id, user_id int, start, end time.Time, title string) (models.Poll, error) {
	poll, err := s.polls.UpdatePoll(id, user_id, start, end, title)
	if err != nil {
		log.Fatal(err)
	}
	return poll, nil
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

func (s *service) GetPollItemsByPollID(poll_id int) ([]models.Item, error) {
	items, err := s.polls.GetPollItemsByPollID(poll_id)
	if err != nil {
		log.Fatal(err)
	}
	return items, nil
}

func (s *service) UpdatePollItem(id, poll_id int, value, display string) (models.Item, error) {
	item, err := s.polls.UpdatePollItem(id, poll_id, value, display)
	if err != nil {
		log.Fatal(err)
	}
	return item, nil
}
func (s *service) DeleteItem(id int) error {
	err := s.polls.DeleteItem(id)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *service) CreateResponse(itemId, pollId int, token string) (models.Response, error) {

	//check cache for key
	cache := cache.NewRedisCache(10)
	key := generateCacheKey(pollId, token)
	value, err := cache.Get(key)

	var response models.Response

	if err != nil || value.(bool) {
		return response, err
	}

	//retrieve related poll
	poll, err := s.polls.GetPoll(pollId)
	if err != nil {
		return response, err
	}

	//do check on expiration just to be sure
	//calculate remaining TTL
	exp := getExpiration(poll.End)

	//expired
	if exp <= 0 {
		return  response,err
	}

	//create response
	response, err = s.polls.CreateResponse(itemId, pollId)
	if err != nil {
		log.Fatal(err)
	}

	//update flag
	cache.SetWithTTL(key, true, exp)

	item := make(chan models.Item)
	user := make(chan models.User)

	//retrieve item async
	go func(s *service, c chan models.Item, itemId int) {
		i, _ := s.GetPollItem(itemId)
		c <- i
	}(s, item, itemId)

	//retrieve user
	go func(s *service, c chan models.User, poll models.Poll) {
		u, _ := s.users.Get(poll.UserID)
		c <- u
	}(s, user, poll)

	//can probably run this async
	s.sendResponseNotification(item, user, poll)

	return response, nil
}

func (s *service) sendResponseNotification(i chan models.Item, u chan models.User, poll models.Poll) {

	item := <-i
	user := <-u

	//the following is only temporary
	templateData := struct {
		Title    string
		Response string
	}{
		Title:    poll.Title,
		Response: item.Display,
	}

	cwd, _ := os.Getwd()
	templatePath := filepath.Join(cwd, "/templates/response_notification.html")
	r := email.NewRequest([]string{user.Email}, "notifications@polling-machine.com", "Your poll received a response!!", "Hello, World!")
	err := r.ParseTemplate(templatePath, templateData)
	if err != nil {
		_, er := r.SendEmail()
		if er != nil {
			log.Println(er)
		}
	}
	r.SendEmail()
}

func (s *service) GetResponseCounts(poll_id int) ([]models.ResponseCount, error) {
	counts, err := s.polls.GetResponseCounts(poll_id)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(counts)
	return counts, nil
}

func (s *service) GetResponseToken(pollId int) (string, error) {

	var token string

	poll, err := s.polls.GetPoll(pollId)
	if err != nil {
		return token, err
	}

	//calculate TTL for key
	exp := getExpiration(poll.End)

	//generate unique response token
	token = uuid.NewUUID().String()

	//generate composite cache key
	key := generateCacheKey(pollId, token)

	//store consumption flag with expiration
	cache := cache.NewRedisCache(10)
	err = cache.SetWithTTL(key, false, exp)

	if err != nil {
		return token, err
	}

	return token, nil
}

func getExpiration(exp time.Time) int {
	return int(exp.Sub(time.Now()).Seconds())
}

func generateCacheKey(pollId int, token string) string {
	return fmt.Sprintf("response%v-%v", pollId, token)
}

func NewService() Service {
	polls, err := NewRepo()
	if err != nil {
		log.Fatal(err)
	}

	users, err := users.NewRepo()

	return &service{
		polls: polls,
		users: users,
	}
}
