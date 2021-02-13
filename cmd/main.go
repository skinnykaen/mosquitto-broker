package main

import (
	"database/sql"
	"log"

	Go "github.com/skinnykaen/go.git"
	"github.com/skinnykaen/go.git/package/handler"
)

func main() {
	db, err := sql.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/users")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT id, email FROM users")

	for results.Next() {
		var users User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&users.ID, &users.Email)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(users.Email)
	}

	handlers := new(handler.Handler)
	srv := new(Go.Server)
	if err := srv.Run("3000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
