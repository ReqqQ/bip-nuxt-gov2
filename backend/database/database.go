package database

import "database/sql"

func connectToDB() (db *sql.DB) {
	db, _ = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/bip")
	return db
}
