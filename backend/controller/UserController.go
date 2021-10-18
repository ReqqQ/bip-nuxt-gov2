package controller

import (
	"cookieApi/backend/database"
	"cookieApi/backend/models"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	db := database.GetDB()

	query, queryError := db.Query("SELECT id,name FROM users")
	if queryError != nil {
		panic(queryError.Error())
	}
	spew.Dump(query)
	var users []models.UserModel
	defer query.Close()
	for query.Next() {
		var user models.UserModel
		if err := query.Scan(&user.Id, &user.Name); err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
