package controller

import (
	"cookieApi/backend/services/posts"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPosts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": posts.GetPosts(),
	})
}
