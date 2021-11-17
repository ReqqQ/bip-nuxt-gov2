package posts

import (
	"cookieApi/backend/database"
	"github.com/jmoiron/sqlx"
)

func getGroupPosts(categoryTypes []string) *sqlx.Rows {
	return getFromDB(categoryTypes, getGroupPostsQuery())
}

func getPostsByName(postName string) *sqlx.Rows {
	db := database.GetDB()
	data, queryError := db.Queryx(getGroupPostsByNameQuery(), postName)
	validateData(queryError)

	return data
}

func getAllPosts() *sqlx.Rows {
	db := database.GetDB()
	data, queryError := db.Queryx(getAllPostsQuery())
	validateData(queryError)

	return data
}

func getFromDB(categoryTypes interface{}, query string) *sqlx.Rows {
	db := database.GetDB()
	queryIn, args, _ := sqlx.In(query, categoryTypes)
	data, queryError := db.Queryx(sqlx.Rebind(sqlx.QUESTION, queryIn), args...)
	validateData(queryError)

	return data
}

func getAllPostsQuery() string {
	return "select id, title, category, created_at from posts"
}

func getGroupPostsByNameQuery() string {
	return "select id, title, category, created_at\nfrom posts p\nwhere MATCH(p.title) AGAINST(? IN NATURAL LANGUAGE MODE)"
}

func getGroupPostsQuery() string {
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
