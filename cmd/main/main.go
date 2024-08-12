package main

import (
	"github.com/gin-gonic/gin"
	"github.com/isabellecostawex/ps-tag-onboarding-go/pkg/postgresql"
	"github.com/isabellecostawex/ps-tag-onboarding-go/internal/api/handlers"

)

func main() {
	err:= postgresql.InitDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	
	router := gin.Default()
	router.POST("/save", handlers.SaveUserHandler)
	router.GET("/find/:id", handlers.FindUserHandler)
	router.Run(":8080")
}
