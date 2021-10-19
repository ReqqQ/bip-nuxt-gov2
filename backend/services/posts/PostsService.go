package posts

import (
	"cookieApi/backend/models"
)

var post models.PostModel

func GetPosts() []models.PostModel {
	post.Reset()
	transformQueryDataToModel(getPostsFromDB(), post)

	return post.GetPosts()
}
