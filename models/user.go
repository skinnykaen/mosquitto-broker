package models

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	u "github.com/skinnykaen/go.git/utils"
	"golang.org/x/crypto/bcrypt"
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
	Password string `json:"password"`
}

var jwtKey = []byte("my_secret_key")

func (user *User) Create () (map[string]interface{}) {
	fakeUser := &User{}
	db, err := sql.Open("mysql", "root:skinny@tcp(127.0.0.1:3306)/mqtt_broker")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	row :=  db.QueryRow("SELECT * from users where email=?", user.UserData.Email).Scan(&fakeUser.Id, &fakeUser.UserData.Email, &fakeUser.UserData.Password)
	if(row == sql.ErrNoRows){

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.UserData.Password), bcrypt.DefaultCost)
		fmt.Println(hashedPassword)
		fmt.Println(len(hashedPassword))
		results, err := db.Exec("INSERT INTO users (email, passwordhash) values (?, ?)", user.UserData.Email, hashedPassword)
		log.Print(results)
		if err != nil {
			panic(err.Error())
		}

		tk := Token{UserId: user.Id}
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
		tokenString, _ := token.SignedString(jwtKey)
		user.Token = tokenString
		fmt.Println("Account has been created")
		response := u.Message(true, "Account has been created")
		response["user"] = user
		return response
	}else{
		fmt.Println("User already exists")
		return u.Message(false, "User already exists")
	}
}

//func (user *User) Create () (map[string]interface{}) {
//	var response string = NewUser(user.UserData.Email, user.UserData.Password)
//
//	switch(response){
//	case "success":


//		tk := Token{UserId: user.Id}
//		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
//		tokenString, _ := token.SignedString(jwtKey)
//		user.Token = tokenString
//		fmt.Println("Account has been created")
//		response := u.Message(true, "Account has been created")
//		response["user"] = user
//		return response
//	default:
//		fmt.Println("User already exists")
//		return u.Message(false, "User already exists")
//	}
//}
//
//func NewUser (emailForm, passwordForm string) string{
//	const res1 = "success"
//	const res2  = "email_already_exists"
//	var found bool
//
//	db, err := sql.Open("mysql", "root:skinny@tcp(127.0.0.1:3306)/mqtt_broker")
//
//	row :=  db.QueryRow("SELECT * from users where email=?", emailForm)
//
//	found, _, _ = FoundEmail(emailForm)
//	if(found){
//		return res2
//	}
//
//	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(passwordForm), bcrypt.DefaultCost)
//
//	results, err := db.Exec("INSERT INTO users (email, passwordhash) values (?, ?)", emailForm, hashedPassword)
//	log.Print(results)
//	if err != nil {
//		panic(err.Error())
//	}
//	return res1
//}

func Login (emailForm, passwordForm string) (map[string]interface{}) {
	user := &User{}

	db, err := sql.Open("mysql", "root:skinny@tcp(127.0.0.1:3306)/mqtt_broker")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	row :=  db.QueryRow("SELECT * from users where email=?", emailForm).Scan(&user.Id, &user.UserData.Email, &user.UserData.Password)
	switch {
	case row == sql.ErrNoRows:
		resp := u.Message(false, "No user with this email")
		return resp
	case row != nil:
		resp := u.Message(false, "Query error")
		return resp
	default:
		err = bcrypt.CompareHashAndPassword([]byte(user.UserData.Password), []byte(passwordForm))
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
			return u.Message(false, "Invalid login credentials. Please try again")
		}
		user.UserData.Email = emailForm
		user.UserData.Password= ""

		tk := &Token{UserId: user.Id}
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
		tokenString, _ := token.SignedString(jwtKey)
		user.Token = tokenString // Сохраните токен в ответе

		fmt.Println("Logged In")
		resp := u.Message(true, "Logged In")
		resp["user"] = user
		return resp
	}
}

//func Login (email, password string) (map[string]interface{}) {
//
//	user := &User{}
//	user.UserData.Email = email
//	user.UserData.Password = password
//
//	switch CheckEmail(user) {
//	case res1:
//		user.UserData.Email = email
//		user.UserData.Password= ""
//
//		tk := &Token{UserId: user.Id}
//		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
//		tokenString, _ := token.SignedString(jwtKey)
//		user.Token = tokenString // Сохраните токен в ответе
//
//		fmt.Println("Logged In")
//		resp := u.Message(true, "Logged In")
//		resp["user"] = user
//		return resp
//	default:
//		fmt.Println("!Logged In")
//		return u.Message(false, "Invalid login credentials. Please try again")
//	}
//
//}
//
//func CheckEmail (user *User) string{
//	var found bool
//	var passwordhashDb string
//	var userId uint
//
//	found, userId, passwordhashDb = FoundEmail(user.UserData.Email)
//	if(found){
//		if(passwordhashDb == user.UserData.Password) {
//			user.Id = userId
//			return res1
//		}else {
//			return res3
//		}
//	}else {
//		return res2
//	}
//	return "error"
//}
//
//func FoundEmail (emailForm string) (bool, uint, string) {
//	db, err := sql.Open("mysql", "root:skinny@tcp(127.0.0.1:3306)/mqtt_broker")
//	if err != nil {
//		panic(err.Error())
//	}
//	defer db.Close()
//
//	results, err := db.Query("SELECT * from users ")
//	defer results.Close()
//	var user User
//
//	for (results.Next()) {
//		err = results.Scan(&user.Id, &user.UserData.Email, &user.UserData.Password)
//		if(user.UserData.Email == emailForm){
//			return true, user.Id, user.UserData.Password
//		}
//	}
//	return false, 0, ""
//}

func (user *User) GetsUserInfo(id uint) (map[string]interface{}) {
	db, err := sql.Open("mysql", "root:skinny@tcp(127.0.0.1:3306)/mqtt_broker")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	row :=  db.QueryRow("SELECT * from users where id=?", id).Scan(&user.Id, &user.UserData.Email, &user.UserData.Password)
	fmt.Println(row)
	
	defer db.Close()

	resp := u.Message(true, "GetsUserInfo success")
	resp["user"] = user
	return resp
}