package service

import (
	"github.com/skinnykaen/mqtt-broker"
	"github.com/skinnykaen/mqtt-broker/package/repository"
)

type TopicsService struct {
	repo repository.Topics
}

func NewTopicsService (repo *repository.Repository) *TopicsService {
	return &TopicsService{repo: repo}
}

func (s *TopicsService) CreateTopic(topic mqtt.Topic) (uint, error) {
	if(topic.TopicData.Password == ""){
		password, err := s.repo.GetUserPassword(topic.Id_User)
		if err != nil {
			return 0, err
		}
		topic.TopicData.Password = password
	}else {
		topic.TopicData.Password = generatePasswordHash(topic.TopicData.Password)
	}
	return  s.repo.CreateTopic(topic)
}

func (s *TopicsService) GetTopics(id uint) ([]*mqtt.Topic, error) {
	return s.repo.GetTopics(id)
}

func (s *TopicsService) Delete (idTopic int) error {
	return s.repo.Delete(idTopic)
}
