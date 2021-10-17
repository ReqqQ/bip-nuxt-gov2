package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/bip")
	if err != nil {
		panic(err.Error())
	}

	return db, err
}
