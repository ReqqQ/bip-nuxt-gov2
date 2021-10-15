package models

import (
	"cookieApi/backend/database"
)

type UserDbRepository interface {
	FindById(id int)
}

func FindById() {
	_ = database.Query("SELECT * FROM Employee ORDER BY id DESC")
}
