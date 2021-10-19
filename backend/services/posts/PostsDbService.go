package posts

import (
	"cookieApi/backend/database"
	"database/sql"
)

func getPostsFromDB() *sql.Rows {
	db := database.GetDB()
	query, queryError := db.Query(getPostsQuery())
	validateData(queryError)

	return query
}

func getPostsQuery() string {
	return "SELECT id,title FROM posts"
}
