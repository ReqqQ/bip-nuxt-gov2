package posts

import (
	"cookieApi/backend/models"
	"github.com/jmoiron/sqlx"
)

var post models.PostModel

func GetPosts() []models.PostModel {
	post.Reset()
	transformQueryDataToModel(getPostsFromDB())

	return post.GetPosts()
}

func transformQueryDataToModel(query *sqlx.Rows) {
	for query.Next() {
		validateData(getQueryScan(query))
		models.AddPost(post)
	}
	query.Close()
}

func getQueryScan(query *sqlx.Rows) error {
	return query.StructScan(&post)
}

func validateData(queryError error) {
	if !isNil(queryError) {
		getError(queryError)
	}
}

func isNil(queryError error) bool {
	return queryError == nil
}

func getError(queryError error) {
	panic(queryError.Error())
}
