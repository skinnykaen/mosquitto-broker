package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/skinnykaen/go.git/models"
	u "github.com/skinnykaen/go.git/utils"

	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		fmt.Println("Invalid request")
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := user.Create() //Создать аккаунт
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(user) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(user.UserData.Email, user.UserData.PasswordHash)
	u.Respond(w, resp)
}