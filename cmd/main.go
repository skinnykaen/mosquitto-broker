package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"log"

	"github.com/skinnykaen/go.git/package/handler"
	"github.com/skinnykaen/go.git/repository"
	"github.com/skinnykaen/go.git/service"

	Go "github.com/skinnykaen/go.git"
)

func main() {
	db, err := sql.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/users")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT id, email FROM users")
	var users Go.User

	for results.Next() {

		err = results.Scan(&users.Id, &users.Email)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		log.Printf(users.Email)
	}

	err = db.QueryRow("SELECT id, email FROM users where id = ?", 1).Scan(&users.Id, &users.Email)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	log.Println(users.Id)
	log.Println(users.Email)

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(Go.Server)
	if err := srv.Run("3000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
