package repository

import (
	"database/sql"
	"fmt"
	"github.com/skinnykaen/mqtt-broker"
)

type AuthMysql struct {
	db *sql.DB
}

func NewAuthMysql(db *sql.DB) *AuthMysql {
	return &AuthMysql{db: db}
}

func (r *AuthMysql) CreateUser(user mqtt.User) (error) {
	query := fmt.Sprintf("INSERT INTO %s (Email, Password) values (?, ?)", usersTable)
	_, err := r.db.Exec(query, user.UserData.Email, user.UserData.Password)
	 if err != nil {
		return  err
	}
	return nil
}

func (r *AuthMysql) GetUser(email, password string) (mqtt.User, error) {
	var user mqtt.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE Email=? AND PASSWORD=?", usersTable)
	row := r.db.QueryRow(query, email, password)
	err := row.Scan(&user.Id, &user.UserData.FirstName, &user.UserData.LastName, &user.UserData.Email, &user.UserData.Password, &user.UserData.MossquittoOn)
	return user, err
}

func (r *AuthMysql) FindUser(email string) (bool) {
	var i string // bad solution
	query := fmt.Sprintf("SELECT Email FROM %s WHERE Email=?", usersTable)
	row := r.db.QueryRow(query, email).Scan(&i)
	if row == sql.ErrNoRows{
		return false
	}else {
		return true
	}
}