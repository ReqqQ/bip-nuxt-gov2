package posts

import "github.com/gin-gonic/gin"

const emptyString = ""

func validateApiParams(c *gin.Context) (string, []string, bool) {
	categoryPosts, _ := c.GetQueryArray("categoryPosts")
	isGroupByCategory := convertStringToBool(c.GetQuery("isGroupByCategory"))
	searchQuery, _ := c.GetQuery("findPosts")

	return searchQuery, categoryPosts, isGroupByCategory
}

func isEmptyPostName(postName string) bool {
	return postName == emptyString
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
