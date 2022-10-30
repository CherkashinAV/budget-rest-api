package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/CherkashinAV/budget-rest-api/database"
)

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.Run("localhost:8080")
}

func getUsers (c *gin.Context) {
	users, err := database.GetUsers()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, nil)
	}
	c.IndentedJSON(http.StatusOK, users)
	return 
}