package posts

import (
	"cookieApi/backend/database"
	"github.com/jmoiron/sqlx"
)

func getPostsFromDB() *sqlx.Rows {
	db := database.GetDB()
	query, queryError := db.Queryx(getPostsQuery())
	validateData(queryError)

	return query
}

func getPostsQuery() string {
	return "select * from ( " +
		"select " +
		"id, " +
		"title, " +
		"category, " +
		"row_number() over (partition by category order by created_at desc) as category_rank from posts p " +
		") ranks " +
		"where category_rank <= 5"
}
