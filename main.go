package main

import (
	backend
)

func main() {
	server :=
	// Set the router as the default one shipped with Gin
//	router := gin.Default()

	// Serve frontend static files
//	router.Use(static.Serve("/", static.LocalFile("./frontend", true)))
	//fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	//db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/")

	// if there is an error opening the connection, handle it
	//if err != nil {
	//	panic(err.Error())
	//}

	// defer the close till after the main function has finished
	// executing
//	defer db.Close()
	// Setup route group for the API
	//api := router.Group("/api")
	//{
	//	api.GET("/", func(c *gin.Context) {
	//		c.JSON(http.StatusOK, gin.H {
	//			"message": "pong",
	//		})
	//	})
	//}

	// Start and run the server
	//router.Run(":3000")
}