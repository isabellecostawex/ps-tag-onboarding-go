package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserData struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

var users = map[string]UserData{}
var currentID int

func generateID() string {
	currentID++
	return strconv.Itoa(currentID)
}

func saveUser(c *gin.Context) {
	var newUser UserData

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User did not pass validation"})
		return
	}

	if newUser.Age < 18 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User did not pass validation", "details": "User does not meet minimum age requirement"})
		return
	}

	if newUser.FirstName == "" || newUser.LastName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User did not pass validation", "details": "User first/last names required"})
		return
	}

	if newUser.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User did not pass validation", "details": "User email required"})
		return
	}

	if !strings.Contains(newUser.Email, "@") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User did not pass validation", "details": "User email must be properly formatted"})
		return
	}

	if newUser.ID != "" {
		_, exists := users[newUser.ID]
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User Not Found"})
			return
		}
		users[newUser.ID] = newUser
		c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "updated_user": newUser})
		return
	}

	newUser.ID = generateID()
	users[newUser.ID] = newUser
	c.JSON(http.StatusCreated, newUser)
}

func findUser(c *gin.Context) {
	userID := c.Param("id")

	user, exists := users[userID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User Not Found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func main() {
	router := gin.Default()
	router.POST("/save", saveUser)
	router.GET("/find/:id", findUser)
	router.Run("localhost:8080")
}
