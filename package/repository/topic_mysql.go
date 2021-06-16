package repository

import (
	"database/sql"
	"fmt"
	"github.com/skinnykaen/mqtt-broker"
)

type TopicsMysql struct {
	db *sql.DB
}

func NewTopicsMysql(db *sql.DB) *TopicsMysql {
	return &TopicsMysql{db: db}
}

func (r *TopicsMysql) CreateTopic (topic mqtt.Topic) (uint, error) {
	var id uint
	query := fmt.Sprintf("INSERT INTO %s (User_Id, TopicName, Password) values (?, ?, ?)", topicsTable)
	_, err := r.db.Exec(query, topic.Id_User , topic.TopicData.Name, topic.TopicData.Password)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *TopicsMysql) GetTopics (id uint) ([]*mqtt.Topic, error) {
	topics := make([]*mqtt.Topic, 0)
	query := fmt.Sprintf("SELECT * FROM %s WHERE User_Id=?", topicsTable)

	results, err := r.db.Query(query, id)
	if err != nil {
		panic(err.Error())
	}

	for (results.Next()) {
		t := &mqtt.Topic{}
		err = results.Scan(&t.Id, &t.Id_User, &t.TopicData.Name, &t.TopicData.Password)
		topics = append(topics, t)
	}
	return topics, err
}

func (r *TopicsMysql) GetUserPassword (id uint) (string, error) {
	var userPassword string
	query := fmt.Sprintf("SELECT Password FROM %s WHERE User_Id=? ", usersTable)
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&userPassword); err != nil {
		return "", err
	}
	return userPassword, nil
}

func (r *TopicsMysql) Delete (idTopic int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE Topic_Id=?", topicsTable)
	_, err := r.db.Exec(query, idTopic)
	if err != nil {
		return err
	}
	return nil
}