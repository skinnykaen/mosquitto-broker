package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/skinnykaen/go.git/models"
	u "github.com/skinnykaen/go.git/utils"
	"net/http"
)

var DeleteTopic = func(w http.ResponseWriter, r *http.Request) {
	topic := &models.Topic{}
	err := json.NewDecoder(r.Body).Decode(topic)
	fmt.Println(topic.TopicData.Name)
	if err != nil {
		fmt.Println("Invalid request")
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := topic.Delete()
	//mosquitto.DeleteFromAclFile("rupychman@mail.ru")
	u.Respond(w, resp)
}

var CreateTopic = func(w http.ResponseWriter, r *http.Request) {
	topic := &models.Topic{}
	err := json.NewDecoder(r.Body).Decode(topic)
	topic.Id_User = r.Context().Value("user").(uint)
	if err != nil {
		fmt.Println("Invalid request")
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := topic.Create()
	u.Respond(w, resp)
}

var GetListTopics = func(w http.ResponseWriter, r *http.Request) {
	topic := &models.Topic{}
	id := r.Context().Value("user").(uint)
	data := topic.GetList(id)
	resp := u.Message(true, "succes");
	resp["data"] = data
	u.Respond(w, resp)
}