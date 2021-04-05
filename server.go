package Go

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/skinnykaen/go.git/app"
	"github.com/skinnykaen/go.git/controllers"
	"net/http"
)

type Server struct {
	*mux.Router
}

func (s *Server) routes (){
	s.Use(app.JwtAuthentication)

	s.HandleFunc("/registration",
		controllers.CreateAccount).Methods("POST")

	s.HandleFunc("/login",
		controllers.Authenticate).Methods("POST")
}

func NewServer() http.Handler {
	s := &Server{
		Router : mux.NewRouter(),
	}
	s.routes()

	handler := cors.New(cors.Options{AllowedOrigins: []string{"http://127.0.0.1", "http://localhost:3000"}}).Handler(s)
	return handler
}

