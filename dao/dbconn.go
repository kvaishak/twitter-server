package dao

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DbConn() (db *sql.DB) {

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/twitterdb2")
	if err != nil {
		panic(err.Error())
	}
	return db
}
