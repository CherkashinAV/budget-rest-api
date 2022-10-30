package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db *DataBase

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.Run("localhost:8080")
	db := Connect()
}

func getUsers (c *gin.Context) {
	var users []User
	var err error
	users, err = db.getUsers()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, nil)
	}
	c.IndentedJSON(http.StatusOK, users)
	return 
}