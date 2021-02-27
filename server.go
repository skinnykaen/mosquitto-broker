package Go

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type Server struct {
	*mux.Router
}

func LoginHandler (w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	log.Println(user.UserData.Email, user.Id)
	if err != nil {
		panic(err.Error())
	}
	email :=user.UserData.Email
	if(checkemail(email)) {
		json.NewEncoder(w).Encode("hello")
	}else {
		json.NewEncoder(w).Encode("no enter")
	}
}

func (s *Server) routes (){
	//s.HandleFunc("/", IndexPageHandler)
	//s.HandleFunc("/internal", InternalPageHandler)

	s.HandleFunc("/login", LoginHandler).Methods("POST")
	//s.HandleFunc("/logout", LogoutHandler).Methods("POST")
}

func NewServer() http.Handler {
	s := &Server{
		Router : mux.NewRouter(),
	}
	handler := cors.New(cors.Options{AllowedOrigins: []string{"http://127.0.0.1", "http://localhost:3000"}}).Handler(s)
	return handler
}

