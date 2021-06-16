package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	usersTable = "users"
	topicsTable = "topics"
)

type Config struct {
	Host string
	Port string
	Username string
	Password string
	DBName string
	SSLMode string
}

func NewMysqlDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err !=nil {
		return nil, err
	}
	return db, nil
}