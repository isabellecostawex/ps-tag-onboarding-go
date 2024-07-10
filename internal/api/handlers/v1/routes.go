package v1

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, handler *UserHandler) {
	router.POST("/save", handler.SaveUserHandler)
	router.GET("/find/:id", handler.FindUserHandler)
}
