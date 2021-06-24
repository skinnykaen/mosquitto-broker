package repository

import (
	"database/sql"
	"fmt"
)

type MosquittoMysql struct {
	db *sql.DB
}

func NewMosquittoMysql(db *sql.DB) *MosquittoMysql {
	return &MosquittoMysql{db: db}
}

func (r *MosquittoMysql) SetMosquittoOn(id uint) (error) {
	query := fmt.Sprintf("UPDATE %s SET MosquittoOn=1 WHERE User_Id=?", usersTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return  err
	}
	return nil
}

func (r *MosquittoMysql) SetMosquittoOff(id uint) (error) {
	query := fmt.Sprintf("UPDATE %s SET MosquittoOn=0 WHERE User_Id=?", usersTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return  err
	}
	return nil
}

func (r *MosquittoMysql) GetMosquittoOn(id uint) (bool, error) {
	var mosquittoOn bool
	query := fmt.Sprintf("SELECTE MosquittoOn FROM %s WHERE User_Id=?", usersTable)
	row := r.db.QueryRow(query, id)
	err := row.Scan(&mosquittoOn)
	return mosquittoOn, err
}

