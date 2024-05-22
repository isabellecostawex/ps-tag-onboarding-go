package main

import (
	"fmt"
	"log"
	"net/http"
	
	"github.com/gin-gonic/gin"
	// "database"
	// "handlers"

)

func main() {
	// err:= database.initDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	
	router := gin.Default()
	// router.POST("/save", api.saveUser)
	// router.GET("/find/:id", api.findUser)
	router.Run(":8080")
}
