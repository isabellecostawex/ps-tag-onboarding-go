package handlers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/isabellecostawex/ps-tag-onboarding-go/internal/api/handlers/v1"
)

func RegisterRoutes (router *gin.Engine, handler *v1.UserHandler) {
	v1.RegisterRoutes(router, handler)
}
