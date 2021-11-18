package posts

import (
	"cookieApi/backend/models"
	"github.com/jmoiron/sqlx"
)

var post models.PostModel
var posts *sqlx.Rows

func GroupByCategory(queryParams []string) map[string][]models.PostModel {
	post.Reset()
	transformQueryDataToModel(getPostsByCategory(queryParams))

	return transformPosts(post.GetPosts())
}

func GetPosts(postName string) []models.PostModel {
	post.Reset()
	transformQueryDataToModel(getPosts(postName))

	return post.GetPosts()
}

func getPosts(postName string) *sqlx.Rows {
	if isEmptyPostName(postName) {
		posts = getAllPosts()
	} else {
		posts = getPostsByName(postName)
	}

	return posts
}

func getQueryScan(query *sqlx.Rows) error {
	return query.StructScan(&post)
}
