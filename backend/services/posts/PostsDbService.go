package posts

import (
	"cookieApi/backend/database"
	"github.com/jmoiron/sqlx"
)

func getPostsByCategory(categoryTypes []string) *sqlx.Rows {
	db := database.GetDB()
	queryIn, args, _ := sqlx.In(getGroupPostsQuery(), categoryTypes)
	data, queryError := db.Queryx(db.Rebind(queryIn), args...)
	validateData(queryError)

	return data
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

func getAllPostsQuery() string {
	return `select id, title, category, created_at 
			from posts`
}

func getGroupPostsByNameQuery() string {
	return `select id, title, category, created_at
			from posts p
			where match(p.title) against(? in natural language mode)`
}

func getGroupPostsQuery() string {
	return `select *
			from ( select id,
						  title,
						  category,
						  row_number() over (partition by category order by created_at desc) as category_rank,
						  created_at
				   from posts p
				   where p.category in (?) ) ranks
			where category_rank <= 6`
}
