package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	Go "github.com/skinnykaen/go.git"
	"log"
	"net/http"
	//"encoding/json"
)

var rooter *mux.Router

func LoginHandler (w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)
	var email = vars["email"]
	//хочу вывести vars
	for k, v := range vars {
		log.Println("key: %d, value: %t\n", k, v)
	}

	if(checkemail(email)) {
		fmt.Fprint(w, "hello")
	}else {
		fmt.Fprint(w,"no enter")
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
		log.Println(email) // из формы
		// получается false потому что email пустая
		if(user.Email == email) {
			return true
		}
	}
	return false
}

func main() {
	rooter := mux.NewRouter()
	rooter.HandleFunc("/hello", LoginHandler).Methods("GET")
	http.Handle("/", rooter)
	handler := cors.New(cors.Options{AllowedOrigins: []string{"http://127.0.0.1", "http://localhost:3000"}}).Handler(rooter)
	log.Fatal(http.ListenAndServe(":8000", handler))

}