package models

import (
	"database/sql"
	u "github.com/skinnykaen/go.git/utils"
	"fmt"
)

type Topic struct {
	Id uint `json:"id"`
	Id_User uint `json:"id_user"`
	TopicData TopicData `json:"topic_data"`
}

type TopicData struct {
	Name string `json:"topicname"`
	Password string `json:"passwordtopic"`
	Mqqt_Tcp_Port uint `json:"mqqt_tcp_port"`
	Secure_mqtt uint `json:"secure_mqtt"`
	Mqtt_Over_Websocket_Port uint `json:"mqtt_over_websocket_port"`
}

func (topic *Topic) Create () (map[string]interface{}){
	db, err := sql.Open("mysql", "root:skinny@tcp(127.0.0.1:3306)/mqtt_broker")

	topics, err := db.Exec("INSERT INTO topics (name, password, mqtt_tcp_port, secure_mqtt, mqtt_over_websocket_port, id_user) values (?,?,?,?,?,?)", topic.TopicData.Name, topic.TopicData.Password , 1883, 3883, 8883, topic.Id_User)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(topics)
	response := u.Message(true, "Topic has been created")
	response["topic"] = topic
	return response
}

// func (topic *Topic) Delete (id_user uint, id_topic uint) (map[string]interface{}){
// 	db, err := sql.Open("mysql", "root:skinny@tcp(127.0.0.1:3306)/mqtt_broker")

// }

func (topic *Topic) GetList (id_user uint) ([]*Topic) {
	db, err := sql.Open("mysql", "root:skinny@tcp(127.0.0.1:3306)/mqtt_broker")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	topics := make([]*Topic, 0)

	results, err := db.Query("SELECT * from topics where id_user=?", id_user)
	defer results.Close()

	for (results.Next()) {
		t := &Topic{}
		err = results.Scan(&t.Id, &t.TopicData.Name, &t.TopicData.Password, &t.TopicData.Mqqt_Tcp_Port, &t.TopicData.Secure_mqtt, &t.TopicData.Mqtt_Over_Websocket_Port, &t.Id_User)
		topics = append(topics, t)
	}	
	return topics
}