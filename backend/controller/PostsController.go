package controller

import (
	"cookieApi/backend/services/posts"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPosts(c *gin.Context) {
	categoryPosts, _ := c.GetQueryArray("categoryPosts")
	isGroupByCategory := posts.ConvertStringToBool(c.GetQuery("isGroupByCategory"))

	data := posts.GetPosts(categoryPosts)
	if isGroupByCategory {
		data = nil
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
