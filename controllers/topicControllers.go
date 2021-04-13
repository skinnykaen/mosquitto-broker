package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/skinnykaen/go.git/models"
	u "github.com/skinnykaen/go.git/utils"
	"net/http"
)

var CreateTopic = func(w http.ResponseWriter, r *http.Request) {
	topic := &models.Topic{}
	err := json.NewDecoder(r.Body).Decode(topic)

	if err != nil {
		fmt.Println("Invalid request")
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := topic.Create(topic.Id_User, topic.Name)
	u.Respond(w, resp)
}

var GetListTopic = func(w http.ResponseWriter, r *http.Request) {
	topic := &models.Topic{}
	id := r.Context().Value("user").(uint)
	resp := topic.GetList(id);
	u.Respond(w, resp)
}