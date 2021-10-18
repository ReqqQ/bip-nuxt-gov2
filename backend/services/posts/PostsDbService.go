package posts

import (
	"cookieApi/backend/database"
	"database/sql"
)

func GetPostsFromDB() *sql.Rows {
	db := database.GetDB()
	query, queryError := db.Query("SELECT id,title FROM posts")
	if queryError != nil {
		panic(queryError.Error())
	}

	return query
}
