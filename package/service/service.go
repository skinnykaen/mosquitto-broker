package service

import (
	"github.com/skinnykaen/mqtt-broker"
	"github.com/skinnykaen/mqtt-broker/package/repository"
)

type Authorization interface {
	CreateUser(user mqtt.User) (error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (uint, error)
}

type Profile interface {
	GetProfile(id uint)(mqtt.User, error)
	SetMosquittoOn(id uint) (error)
	SetMosquittoOff(id uint) (error)
}

type Topics interface {
	CreateTopic(topic mqtt.Topic) (uint, error)
	GetTopics(id uint) ([]*mqtt.Topic, error)
	Delete(idTopic int) error
}

type Service struct {
	Authorization
	Profile
	Topics
}
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		Profile: NewProfileService(repos),
		Topics: NewTopicsService(repos),
	}
}