package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	Go "github.com/skinnykaen/go.git"
	"log"
	"net/http"
)

var rooter *mux.Router

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
		err = results.Scan(&user.Id, &user.UserData.Email)
		if(user.UserData.Email == email) {
			return true
		}
	}
	return false
}

func main() {
	//rooter := mux.NewRouter()
	//rooter.HandleFunc("/hello", LoginHandler).Methods("POST")
	//handler := cors.New(cors.Options{AllowedOrigins: []string{"http://127.0.0.1", "http://localhost:3000"}}).Handler(rooter)
	//http.ListenAndServe(":8000", handler)

	srv := Go.NewServer()
	log.Fatal(http.ListenAndServe(":8000", srv))

}