package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	Go "github.com/skinnykaen/go.git"
	"log"
	"net/http"
)

var rooter *mux.Router

func LoginHandler (w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var user Go.User
	err := json.NewDecoder(r.Body).Decode(&user)
	log.Println(user.Email, user.Id)
	if err != nil {
		panic(err.Error())
	}
	email :=user.Email
	if(checkemail(email)) {
		json.NewEncoder(w).Encode("hello")
	}else {
		json.NewEncoder(w).Encode("suck my pipe")
	}
}

func checkemail (email string) bool{
	db, err := sql.Open("mysql", "root:skinny@tcp(127.0.0.1:3306)/mqtt_broker")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT * from users ")
	defer results.Close()
	for (results.Next()) {
		var user Go.User
		err = results.Scan(&user.Id, &user.Email)
		log.Println(user.Email) // из базы данных где лежит один email tortancs@mail.ru

		// получается false потому что email пустая
		if(user.Email == email) {
			return true
		}
	}
	return false
}

func main() {
	rooter := mux.NewRouter()
	rooter.HandleFunc("/hello", LoginHandler).Methods("POST")
	handler := cors.New(cors.Options{AllowedOrigins: []string{"http://127.0.0.1", "http://localhost:3000"}}).Handler(rooter)
	http.ListenAndServe(":8000", handler)

	//srv := Go.NewServer()
	//log.Fatal(http.ListenAndServe(":8000", srv))

}