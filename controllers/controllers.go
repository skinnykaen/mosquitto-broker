package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/skinnykaen/go.git/models"
	u "github.com/skinnykaen/go.git/utils"
	"net/http"
)

var Me = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("wsup")
	id := r.Context().Value("user").(uint)
	fmt.Println(id)
	resp := u.Message(true, "success")
	w.Header().Set("Access-Control-Allow-Origin", "*")
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

	resp := models.Login(user.UserData.Email, user.UserData.PasswordHash)
	u.Respond(w, resp)
}