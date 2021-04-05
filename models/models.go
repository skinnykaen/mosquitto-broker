package models

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	u "github.com/skinnykaen/go.git/utils"
	"log"
)

const res1 = "success"
const res2  = "not_found_email"
const res3 = "invalid_password"

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type User struct {
	Id uint `json:"id"`
	UserData UserData `json:"user_data"`
	Token string `json:"token";sql:"-"`
}

type UserData struct {
	Email string `json:"email"`
	PasswordHash string `json:"passwordhash"`
}

var jwtKey = []byte("my_secret_key")

func (user *User) Create () (map[string]interface{}) {
	var response string = NewUser(user.UserData.Email, user.UserData.PasswordHash)

	switch(response){
	case "success":
		tk := Token{UserId: user.Id}
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
		tokenString, _ := token.SignedString(jwtKey)
		user.Token = tokenString
		fmt.Println("Account has been created")
		response := u.Message(true, "Account has been created")
		response["user"] = user
		return response
	default:
		fmt.Println("User already exists")
		return u.Message(false, "User already exists")
	}
}

func NewUser (emailForm, passwordForm string) string{
	const res1 = "success"
	const res2  = "email_already_exists"
	var found bool

	db, err := sql.Open("mysql", "root:skinny@tcp(127.0.0.1:3306)/mqtt_broker")

	found, _, _ = FoundEmail(emailForm)
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

func Login (email, password string) (map[string]interface{}) {

	user := &User{}
	user.UserData.Email = email
	user.UserData.PasswordHash = password

	switch CheckEmail(user) {
	case res1:
		{
			//expirationTime := time.Now().Add(5 * time.Minute)
			//claims := &models.Token{
			//	Username: user.UserData.Email,
			//	StandardClaims: jwt.StandardClaims{
			//		ExpiresAt: expirationTime.Unix(),
			//	},
			//}
			//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			//tokenString, err := token.SignedString(jwtKey)
			//if err != nil {
			//	w.WriteHeader(http.StatusInternalServerError)
			//	return
			//}
			//http.SetCookie(w, &http.Cookie{
			//	Name:    "token",
			//	Value:   tokenString,
			//	Expires: expirationTime,
			//})
			//Welcome(w,r)
		}

		user.UserData.Email = email
		user.UserData.PasswordHash = ""

		tk := &Token{UserId: user.Id}
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
		tokenString, _ := token.SignedString(jwtKey)
		user.Token = tokenString // Сохраните токен в ответе
		fmt.Println("Logged In")
		resp := u.Message(true, "Logged In")
		resp["user"] = user
		return resp
	default:
		fmt.Println("!Logged In")
		return u.Message(false, "Invalid login credentials. Please try again")
	}

}

func CheckEmail (user *User) string{
	var found bool
	var passwordhashDb string
	var userId uint

	found, userId, passwordhashDb = FoundEmail(user.UserData.Email)
	if(found){
		if(passwordhashDb == user.UserData.PasswordHash) {
			user.Id = userId
			return res1
		}else {
			return res3
		}
	}else {
		return res2
	}
	return "error"
}

func FoundEmail (emailForm string) (bool, uint, string) {
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
			return true, user.Id, user.UserData.PasswordHash
		}
	}
	return false, 0, ""
}