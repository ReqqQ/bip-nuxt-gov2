package posts

import (
	"cookieApi/backend/models"
	"github.com/jmoiron/sqlx"
)

var post models.PostModel

func GroupByCategory(queryParams []string) map[string][]models.PostModel {
	post.Reset()
	transformQueryDataToModel(getGroupPosts(queryParams))

	return transformPosts(post.GetPosts())
}

func GetPosts(postName string) []models.PostModel {
	post.Reset()
	var posts *sqlx.Rows
	if postName == "" {
		posts = getAllPosts()
	} else {
		posts = getPostsByName(postName)
	}
	transformQueryDataToModel(posts)
	return post.GetPosts()
}

func getQueryScan(query *sqlx.Rows) error {
	return query.StructScan(&post)
}
