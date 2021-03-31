package main

import (
	_ "github.com/go-sql-driver/mysql"
	Go "github.com/skinnykaen/go.git"
	"log"
	"net/http"
)

func main() {
	srv := Go.NewServer()

	log.Fatal(http.ListenAndServe(":8000", srv))

}