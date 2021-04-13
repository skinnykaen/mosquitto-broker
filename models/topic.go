package models

import (
	"database/sql"
	"fmt"
	u "github.com/skinnykaen/go.git/utils"
	"log"
)

type Topic struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	Mqqt_Tcp_Port uint `json:"mqqt_tcp_port"`
	Secure_mqtt uint `json:"secure_mqtt "`
	Mqtt_Over_Websocket_Port uint `json:"mqtt_over_websocket_port"`
	Id_User uint `json:"id_user"`
}

func (topic *Topic) Create (id_user uint, name string) (map[string]interface{}){
	user := &User{}

	db, err := sql.Open("mysql", "root:skinny@tcp(127.0.0.1:3306)/mqtt_broker")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	users, err := db.Query("SELECT * from users ")
	defer users.Close()

	for(users.Next()){
		err = users.Scan(&user.Id, &user.UserData.Email, &user.UserData.PasswordHash)
		if(user.Id == id_user){
			break;
		}
		fmt.Println("erorr")
		return u.Message(false, "error")
	}

	name_topic :=  user.UserData.Email + "/" + name

	topics, err := db.Exec("INSERT INTO topics (name, password, mqtt_tcp_port, secure_mqtt,  mqtt_over_websocket_port, id_user) values (?,?,?,?,?,?)", name_topic, user.UserData.PasswordHash, 1883, 3883, 8883, id_user)
	log.Print(topics)
	if err != nil {
		panic(err.Error())
	}
	response := u.Message(true, "Account has been created")
	return response
}

func (topic *Topic) GetList (id_user uint) (map[string]interface{}) {
	db, err := sql.Open("mysql", "root:skinny@tcp(127.0.0.1:3306)/mqtt_broker")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT * from topics ")
	defer results.Close()


	fmt.Println("error in GetList")
	
	for (results.Next()) {
		err = results.Scan(&topic.Id_User, &topic.Name, &topic.Password)
		fmt.Println(topic.Id_User)
		if(topic.Id_User == id_user){
			resp := u.Message(true, "topic exist")
			resp["topic"] = topic
			return resp
		}
	}
	fmt.Println("error")
	return u.Message(false, "error")
}