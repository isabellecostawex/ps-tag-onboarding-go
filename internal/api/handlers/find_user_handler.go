package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/isabellecostawex/ps-tag-onboarding-go/internal/services"
)

func FindUserHandler(c *gin.Context){
	userID := c.Param("id")
	user, err := services.RetrieveUser(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{ "error": "User not found"})
	}
	c.JSON(http.StatusOK, user)
}
