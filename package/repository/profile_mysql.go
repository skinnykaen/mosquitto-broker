package repository

import (
	"database/sql"
	"fmt"
	"github.com/skinnykaen/mqtt-broker"
)

type ProfileMysql struct {
	db *sql.DB
}

func NewProfileMysql (db *sql.DB) *ProfileMysql {
	return &ProfileMysql{db: db}
}

func (r *ProfileMysql) GetProfile(id uint) (mqtt.User, error) {
	var user mqtt.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE User_Id=?", usersTable)
	row := r.db.QueryRow(query, id)
	err := row.Scan(&user.Id, &user.UserData.FirstName, &user.UserData.LastName, &user.UserData.Email, &user.UserData.Password, &user.UserData.MossquittoOn)
	return user, err
}