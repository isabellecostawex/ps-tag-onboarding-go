package main

import (
	"log"
	
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
	handlers.RegisterRoutes(router)
	router.Run(":8080")
}
