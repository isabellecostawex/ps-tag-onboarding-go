package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isabellecostawex/ps-tag-onboarding-go/internal/domain/users/postgres"
	"github.com/isabellecostawex/ps-tag-onboarding-go/internal/services"
)

func FindUserHandler(c *gin.Context) {
	userID := c.Param("id")

	userRepo := &postgres.UsersRepository{}
	usersService := services.UserManagementService{UserRepo: userRepo}

	user, err := usersService.RetrieveUser(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
