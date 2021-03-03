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

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode(NewUser(user.UserData.Email, user.UserData.PasswordHash))
}

func NewUser (emailForm string, passwordForm string) string{
	const res1 = "success"
	const res2  = "email_already_exists"
	var found bool

	db, err := sql.Open("mysql", "root:skinny@tcp(127.0.0.1:3306)/mqtt_broker")

	found, _ = FoundEmail(emailForm)
	if(found){
		return res2
	}

	results, err := db.Exec("INSERT INTO users (email, passwordhash) values (?, ?)", emailForm, passwordForm)
	log.Print(results)
	if err != nil {
		panic(err.Error())
	}
	return res1
}

func LoginHandler (w http.ResponseWriter, r *http.Request)  {
	var user User

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode(CheckEmail(user.UserData.Email, user.UserData.PasswordHash))
}

func CheckEmail (emailForm string, passwordForm string) string{
	const res1 = "success"
	const res2  = "not_found_email"
	const res3 = "invalid_password"

	var found bool
	var passwordhashDb string

	found,passwordhashDb = FoundEmail(emailForm)
	if(found){
		if(passwordhashDb == passwordForm) {
			return res1
		}else {
			return res3
		}
	}else {
		return res2
	}
	return "error"
}

func FoundEmail (emailForm string) (bool, string) {
	db, err := sql.Open("mysql", "root:skinny@tcp(127.0.0.1:3306)/mqtt_broker")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT * from users ")
	defer results.Close()
	var user User

	for (results.Next()) {
		err = results.Scan(&user.Id, &user.UserData.Email, &user.UserData.PasswordHash)
		if(user.UserData.Email == emailForm){
			return true, user.UserData.PasswordHash
		}
	}
	return false, user.UserData.PasswordHash
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

