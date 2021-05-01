package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/skinnykaen/go.git/models"
	"github.com/skinnykaen/go.git/mosquitto"
	u "github.com/skinnykaen/go.git/utils"
	"net/http"
)

var Me = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	id := r.Context().Value("user").(uint)

	resp := user.GetsUserInfo(id);
	u.Respond(w, resp)
}

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		fmt.Println("Invalid request")
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := user.Create() //Создать аккаунт
	args := []string{"-b","passwd.txt", user.UserData.Email, user.UserData.Password}
	mosquitto.RunCommand("mosquitto_passwd", args...) //Записать в passwd
	mosquitto.WriteToAclFile(user.UserData.Email) //Записать в acl
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(user) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	//expirationTime := time.Now().Add(5 * time.Minute)
	//http.SetCookie(w, &http.Cookie{
	//		Name:    "token",
	//		Value:   user.Token,
	//		Expires: expirationTime,
	//})

	resp := models.Login(user.UserData.Email, user.UserData.Password)
	u.Respond(w, resp)
}
