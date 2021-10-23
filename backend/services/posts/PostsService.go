package posts

import (
	"cookieApi/backend/models"
	"github.com/jmoiron/sqlx"
)

var post models.PostModel

func GetPosts(queryParams []string) map[string][]models.PostModel {
	post.Reset()
	transformQueryDataToModel(getPostsFromDB(queryParams))
	posts := post.GetPosts()
	return transformPosts(posts)
}
func transformPosts(posts []models.PostModel) map[string][]models.PostModel {
	transformedArray := make(map[string][]models.PostModel)
	for _, element := range posts {
		transformedArray[element.Category] = append(transformedArray[element.Category], element)
	}

	return transformedArray
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
