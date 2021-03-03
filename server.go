package Go

import (
	"database/sql"
	"encoding/json"
	"log"

	//"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	//"github.com/gorilla/securecookie"

	"net/http"
)

type Server struct {
	*mux.Router
}

//var cookieHandler = securecookie.New(
//securecookie.GenerateRandomKey(64),
//securecookie.GenerateRandomKey(32))
//
//func getUserName(request *http.Request) (userName string) {
//	if cookie, err := request.Cookie("email"); err == nil {
//		cookieValue := make(map[string]string)
//		if err = cookieHandler.Decode("email", cookie.Value, &cookieValue); err == nil {
//			userName = cookieValue["email"]
//		}
//	}
//	return userName
//}
//
//const internalPage = `
//<h1>Internal</h1>
//<hr>
//<small>User: %s</small>
//<form method="post" action="/logout">
//    <button type="submit">Logout</button>
//</form>
//`
//
//func InternalPageHandler(response http.ResponseWriter, request *http.Request) {
//	userName := getUserName(request)
//	if userName != "" {
//		fmt.Fprintf(response, internalPage, userName)
//	} else {
//		http.Redirect(response, request, "/", 302)
//	}
//}

func RegistrationHandler (w http.ResponseWriter, r *http.Request) {
	var user User
	var redirectTarget string = "/"

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err.Error())
	}

	email :=user.UserData.Email
	password := user.UserData.PasswordHash

	if(NewUser(email, password)){
		json.NewEncoder(w).Encode("hello new user")
	}else {
		json.NewEncoder(w).Encode("user exist")
	}
	http.Redirect(w, r, redirectTarget, 302)
}

func NewUser (email string, password string) bool {
	db, err := sql.Open("mysql", "root:skinny@tcp(127.0.0.1:3306)/mqtt_broker")

	if(!checkuser(email, password)){
		return false
	}

	result2, err := db.Exec("INSERT INTO users (email, passwordhash) values (?, ?)", email, password)
	log.Print(result2)
	if err != nil {
		panic(err.Error())
	}
	return true
}

func checkuser(email string, password string) bool{

	db, err := sql.Open("mysql", "root:skinny@tcp(127.0.0.1:3306)/mqtt_broker")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT * from users ")
	defer results.Close()

	for (results.Next()) {
		var user User
		err = results.Scan(&user.Id, &user.UserData.Email, &user.UserData.PasswordHash)
		if(email != "" && password != "") {
			if(user.UserData.Email == email){
				return false
			}
		}
	}
	return true
}

func LoginHandler (w http.ResponseWriter, r *http.Request)  {
	var user User
	var redirectTarget string = "/"

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err.Error())
	}

	email :=user.UserData.Email
	password := user.UserData.PasswordHash

	if(checkemail(email, password)) {
		json.NewEncoder(w).Encode("hello")
		redirectTarget = "google.com"
	}else {
		json.NewEncoder(w).Encode("no enter")
	}
	http.Redirect(w, r, redirectTarget, 302)
}

func checkemail (email string, password string) bool{
	db, err := sql.Open("mysql", "root:skinny@tcp(127.0.0.1:3306)/mqtt_broker")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT * from users ")
	defer results.Close()

	for (results.Next()) {
		var user User
		err = results.Scan(&user.Id, &user.UserData.Email, &user.UserData.PasswordHash)
		if(email != "" && password != "") {
			if(user.UserData.Email == email && user.UserData.PasswordHash == password){
				return true
			}
		}
	}
	return false
}

func (s *Server) routes (){
	//s.HandleFunc("/", IndexPageHandler)
	//s.HandleFunc("/internal", InternalPageHandler)
	s.HandleFunc("/login", LoginHandler).Methods("POST")
	s.HandleFunc("/registration", RegistrationHandler).Methods("POST")
	//s.HandleFunc("/logout", LogoutHandler).Methods("POST")
}

func NewServer() http.Handler {
	s := &Server{
		Router : mux.NewRouter(),
	}
	s.routes()
	handler := cors.New(cors.Options{AllowedOrigins: []string{"http://127.0.0.1", "http://localhost:3000"}}).Handler(s)
	return handler
}

