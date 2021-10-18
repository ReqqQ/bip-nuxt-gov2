package posts

import (
	"cookieApi/backend/models"
	"database/sql"
)

var posts []models.PostModel
var post models.PostModel
var query *sql.Rows

func GetPosts() []models.PostModel {
	query = GetPostsFromDB()
	defer query.Close()
	createPostModel()
	return posts
}

func createPostModel() {
	for query.Next() {
		validateQueryScan()
		posts = append(posts, post)
	}
}

func queryStatus() error {
	return query.Scan(&post.Id, &post.Title)
}

func validateQueryScan() {
	isZeroValue, queryStatus := isZeroValue()
	if !isZeroValue {
		panic(queryStatus.Error())
	}
}

func isZeroValue() (bool, error) {
	status := queryStatus()
	return status == nil, status
}
