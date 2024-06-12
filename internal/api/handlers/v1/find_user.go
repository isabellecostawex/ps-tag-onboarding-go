package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isabellecostawex/ps-tag-onboarding-go/internal/domain/users/postgres"
	"github.com/isabellecostawex/ps-tag-onboarding-go/internal/services"
)

type findUserResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

func FindUserHandler(c *gin.Context) {
	userID := c.Param("id")

	userRepo := &postgres.UsersRepository{}
	usersService := services.UserManagementService{UserRepo: userRepo}

	user, err := usersService.RetrieveUser(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	response := findUserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Age:       user.Age,
	}
	c.JSON(http.StatusOK, response)
}
