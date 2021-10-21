package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var database *sqlx.DB

func ConnectToDB() {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/bip")
	if err != nil {
		panic(err.Error())
	}

	database = db
}
func GetDB() *sqlx.DB {
	return database
}
