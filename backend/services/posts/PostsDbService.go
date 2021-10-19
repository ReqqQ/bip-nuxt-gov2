package posts

import (
	"cookieApi/backend/database"
	"cookieApi/backend/models"
	"database/sql"
)

func getPostsFromDB() *sql.Rows {
	db := database.GetDB()
	query, queryError := db.Query("SELECT id,title FROM posts")
	if queryError != nil {
		panic(queryError.Error())
	}

	return query
}
func transformQueryDataToModel(query *sql.Rows, post models.PostModel) {
	for query.Next() {
		if err := query.Scan(&post.Id, &post.Title); err != nil {
			panic(err.Error())
		}
		models.AddPost(post)
	}
	query.Close()
}
