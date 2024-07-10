package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/isabellecostawex/ps-tag-onboarding-go/internal/api/handlers"
	"github.com/isabellecostawex/ps-tag-onboarding-go/internal/domain/users/postgres"
	"github.com/isabellecostawex/ps-tag-onboarding-go/internal/services"
	"github.com/isabellecostawex/ps-tag-onboarding-go/pkg/postgresql"
)

func main() {
	err := postgresql.InitDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	repo := postgres.UsersRepository{DB: postgresql.DB}
	service := services.UserManagementService{UserRepo: &repo}
	handler := handlers.UserHandler{UserService: service}
	router := gin.Default()
	handlers.RegisterRoutes(router, &handler)

	router.Run(":8080")
}
