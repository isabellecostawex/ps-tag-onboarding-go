package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isabellecostawex/ps-tag-onboarding-go/internal/domain/users"
)

type saveUserRequest struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

type SaveUserResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

func (handler *UserHandler) SaveUserHandler(c *gin.Context) {
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

	userID, err := handler.UserService.SaveUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := SaveUserResponse{
		ID:      userID,
		Message: "User saved successfully",
	}

	c.JSON(http.StatusOK, response)
}
