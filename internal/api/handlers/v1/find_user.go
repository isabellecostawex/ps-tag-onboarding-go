package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FindUserResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

func (handler *UserHandler) FindUserHandler(c *gin.Context) {
	userID := c.Param("id")

	user, err := handler.UserService.RetrieveUser(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	response := FindUserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Age:       user.Age,
	}
	c.JSON(http.StatusOK, response)
}
