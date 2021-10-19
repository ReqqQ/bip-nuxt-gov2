package posts

import (
	"cookieApi/backend/models"
	"database/sql"
)

var post models.PostModel

func GetPosts() []models.PostModel {
	post.Reset()
	transformQueryDataToModel(getPostsFromDB())

	return post.GetPosts()
}

func transformQueryDataToModel(query *sql.Rows) {
	for query.Next() {
		validateData(getQueryScan(query))
		models.AddPost(post)
	}
	query.Close()
}

func getQueryScan(query *sql.Rows) error {
	return query.Scan(&post.ID, &post.Title)
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
