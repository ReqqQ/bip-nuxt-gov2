package posts

import (
	"cookieApi/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"strconv"
)

func GetApiParams(c *gin.Context) (string, []string, bool) {
	return validateApiParams(c)
}

func convertStringToBool(queryKey string, error bool) bool {
	if !error {
		queryKey = string(0)
	}
	convertedKey, _ := strconv.ParseBool(queryKey)
	return convertedKey
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
