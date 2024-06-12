package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isabellecostawex/ps-tag-onboarding-go/internal/domain/users"
	"github.com/isabellecostawex/ps-tag-onboarding-go/internal/domain/users/postgres"
	"github.com/isabellecostawex/ps-tag-onboarding-go/internal/services"
)

type saveUserRequest struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

type saveUserResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

func SaveUserHandler(c *gin.Context) {
	var req saveUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data in request body"})
		return
	}

	newUser := users.UserData{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Age:       req.Age,
	}

	userRepo := &postgres.UsersRepository{}
	usersService := services.UserManagementService{UserRepo: userRepo}

	userID, err := usersService.SaveUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := saveUserResponse{
		ID:      userID,
		Message: "User saved successfully",
	}

	c.JSON(http.StatusOK, response)
}
