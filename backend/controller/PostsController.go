package controller

import (
	"cookieApi/backend/services/posts"
	"github.com/gin-gonic/gin"
	"net/http"
)

var dbPosts interface{}

func GetPosts(c *gin.Context) {
	searchQuery, categoryPosts, isGroupByCategory := posts.GetApiParams(c)

	if isGroupByCategory {
		dbPosts = posts.GroupByCategory(categoryPosts)
	} else {
		dbPosts = posts.GetPosts(searchQuery)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": dbPosts,
	})
}
