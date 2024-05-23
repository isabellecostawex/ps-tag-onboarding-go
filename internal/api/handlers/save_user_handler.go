package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/isabellecostawex/ps-tag-onboarding-go/internal/models"
	"github.com/isabellecostawex/ps-tag-onboarding-go/internal/services"
)

func SaveUserHandler (c *gin.Context) {
	var newUser models.UserData

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data in request body"})
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

	var userID int
	var err error

	if newUser.ID != 0 {
		err = services.UpdateUser(&newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}
		userID = newUser.ID
	} else {
		userID, err = services.CreateUser(&newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
			return
		}
	}
	newUser.ID = userID
	c.JSON(http.StatusOK, gin.H{"message": "User saved successfully", "user": newUser})
}
