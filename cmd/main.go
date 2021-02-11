package main

import (
	"log"

	Go "github.com/skinnykaen/go.git"
	"github.com/skinnykaen/go.git/package/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(Go.Server)
	if err := srv.Run("3000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
