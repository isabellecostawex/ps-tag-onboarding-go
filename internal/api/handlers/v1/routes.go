package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/isabellecostawex/ps-tag-onboarding-go/internal/api/handlers"
)

func RegisterRoutes(router *gin.Engine, handler *handlers.UserHandler) {
	router.POST("/save", handler.SaveUserHandler)
	router.GET("/find/:id", handler.FindUserHandler)
}
