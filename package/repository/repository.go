package repository

import (
	"database/sql"
	"github.com/skinnykaen/mqtt-broker"
)

type Authorization interface {
	CreateUser(user mqtt.User) (error)
	GetUser(email, password string) (mqtt.User, error)
	FindUser(email string) (bool)
}

type Profile interface {
	GetProfile(id uint) (mqtt.User, error)
}

type Topics interface {
	CreateTopic(topic mqtt.Topic) (uint, error)
	GetTopics(id uint) ([]*mqtt.Topic, error)
	GetUserPassword (id uint) (string, error)
	Delete(idTopic int) error
}

type Mosquitto interface {
	SetMosquittoOn(id uint) (error)
	SetMosquittoOff(id uint) (error)
	GetMosquittoOn(id uint) (bool, error)
}

type Repository struct {
	Authorization
	Profile
	Topics
	Mosquitto
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMysql(db),
		Profile: NewProfileMysql(db),
		Topics: NewTopicsMysql(db),
		Mosquitto: NewMosquittoMysql(db),
	}
}
