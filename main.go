package main

import (
	"net/http"

	"github.com/CherkashinAV/budget-rest-api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserById)
	router.GET("/users/:id/transactions", getAllUserTransactions)
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


func getUserById (c * gin.Context) {
	id := c.Param("id")
	user, err := database.GetUserById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, nil)
		return 
	}
	c.IndentedJSON(http.StatusOK, user)
	return
}


func getAllUserTransactions (c *gin.Context) {
	id := c.Param("id")
	tr, err := database.GetAllUserTransactions(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, nil)
		return 
	}
	c.IndentedJSON(http.StatusOK, tr)
	return
}

