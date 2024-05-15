package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type UserData struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:123456@postgres:5432/users_database?sslmode=disable")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(100) NOT NULL,
		last_name VARCHAR(100) NOT NULL,
		email VARCHAR(250) NOT NULL,
		age INT NOT NULL
	)`)

	if err != nil {
		panic(err)
	}

}

func saveUser(c *gin.Context) {
	var newUser UserData

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User did not pass validation"})
		return
	}

	if newUser.Age < 18 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User did not pass validation", "details": "User does not meet minimum age requirement"})
		return
	}

	if newUser.FirstName == "" || newUser.LastName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User did not pass validation", "details": "User first/last names required"})
		return
	}

	if newUser.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User did not pass validation", "details": "User email required"})
		return
	}

	if !strings.Contains(newUser.Email, "@") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User did not pass validation", "details": "User email must be properly formatted"})
		return
	}

	var userID string

	if newUser.ID != "" {
		_, err := db.Exec("UPDATE users SET first_name=$1, last_name=$2, email=$3, age=$4 WHERE id=$5", newUser.FirstName, newUser.LastName, newUser.Email, newUser.Age, newUser.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update User"})
			return
		}
	} else {
		_, err := db.Exec("INSERT INTO users (id, first_name, last_name, email, age) VALUES ($1, $2, $3, $4, $5)", "", newUser.FirstName, newUser.LastName, newUser.Email, newUser.Age)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save User"})
			fmt.Println(err)
			return
		}

	}

	newUser.ID = userID
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "updated_user": newUser})
}

func findUser(c *gin.Context) {
	userID := c.Param("id")
	var user UserData
	err := db.QueryRow("SELECT id, first_name, last_name, email, age FROM users WHERE ID=$1", userID).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Age)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User Not Found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func main() {
	initDB()

	router := gin.Default()
	router.POST("/save", saveUser)
	router.GET("/find/:id", findUser)
	router.Run("localhost:8080")
}
