package database

import "github.com/jmoiron/sqlx"

type Database struct {
	Ip       string
	Port     int
	Name     string
	User     string
	Password string

	Index  int
	DB     *sqlx.DB
	Config string
}
