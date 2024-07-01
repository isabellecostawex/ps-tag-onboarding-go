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
	
	repo := user.UsersRepository{DB: postgresql.DB}
	service := services.UserManagementService{UserRepo: &repo}
	handler := handlers.UserHandler{UserService: service}
	router := handlers.RegisterRoutes(&handler)

	router.Run(":8080")
}
