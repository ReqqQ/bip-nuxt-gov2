package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

func ConnectToDB() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/bip")
	if err != nil {
		panic(err.Error())
	}

	database = db
}
func GetDB() *sql.DB {
	return database
}
