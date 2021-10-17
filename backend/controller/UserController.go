package controller

import (
	"cookieApi/backend/database"
	"cookieApi/backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {
	db, _ := database.ConnectToDB()

	query, queryError := db.Query("SELECT id,name FROM users")
	if queryError != nil {
		panic(queryError.Error())
	}

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
