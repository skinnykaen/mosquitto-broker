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

func (r *ProfileMysql) SetMosquittoOn(id uint) (error) {
	query := fmt.Sprintf("UPDATE %s SET MosquittoOn=1 WHERE User_Id=?", usersTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return  err
	}
	return nil
}

func (r *ProfileMysql) SetMosquittoOff(id uint) (error) {
	query := fmt.Sprintf("UPDATE %s SET MosquittoOn=0 WHERE User_Id=?", usersTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return  err
	}
	return nil
}