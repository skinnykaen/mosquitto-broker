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