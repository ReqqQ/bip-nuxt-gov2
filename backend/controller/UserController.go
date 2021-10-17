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

	query.Next()
	{
		var user models.UserModel
		err := query.Scan(&user.Id, &user.Name)
		if err != nil {
			return
		} else {
			users = append(users, user)
		}
	}
	query.Close()
	db.Close()
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
