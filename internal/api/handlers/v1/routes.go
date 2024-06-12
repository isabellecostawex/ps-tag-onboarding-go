package handlers

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	router.POST("/save", SaveUserHandler)
	router.GET("find/:id", FindUserHandler)
}
