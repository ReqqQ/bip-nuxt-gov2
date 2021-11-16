package posts

import (
	"cookieApi/backend/database"
	"github.com/jmoiron/sqlx"
)

func getPostsFromDB(categoryTypes []string) *sqlx.Rows {
	db := database.GetDB()
	queryIn, args, _ := sqlx.In(getPostsQuery(), categoryTypes)
	data, queryError := db.Queryx(sqlx.Rebind(sqlx.QUESTION, queryIn), args...)
	validateData(queryError)

	return data
}

func getPostsQuery() string {
	return "select * from ( " +
		"select " +
		"id, " +
		"title, " +
		"category, " +
		"row_number() over (partition by category order by created_at desc) as category_rank, " +
		"created_at  from posts p where p.category IN(?)" +
		") ranks " +
		"where category_rank <= 6"
}
