package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/isabellecostawex/ps-tag-onboarding-go/internal/api/handlers"
)

func RegisterRoutes (router *gin.Engine) {
	handlers.RegisterRoutes(router)
}
