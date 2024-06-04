package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isabellecostawex/ps-tag-onboarding-go/internal/domain/users"
	"github.com/isabellecostawex/ps-tag-onboarding-go/internal/domain/users/postgres"
	"github.com/isabellecostawex/ps-tag-onboarding-go/internal/services"
)

func SaveUserHandler(c *gin.Context) {
    var newUser users.UserData

    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data in request body"})
        return
    }

    userRepo := &postgres.UsersRepository{}
    usersService := services.UserManagementService{UserRepo: userRepo}

    userID, err := usersService.SaveUser(newUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    newUser.ID = userID
    c.JSON(http.StatusOK, gin.H{"message": "User saved successfully", "user": newUser})
}
