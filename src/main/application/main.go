package main

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type UserData struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

var db *sql.DB
// var users = map[string]UserData{}
// var currentID int

func initDB(){
	var err error
	db, err = sql.Open("postgres", "postgres://user:password@localhost/database_name?sslmode=disable")
	if err !nil {
		panic(err)
	}
}

/*
func generateID() string {
	currentID++
	return strconv.Itoa(currentID)
}
*/

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
		err := db.QueryRow("SELECT id FROM users WHERE id=$1", newUser.ID).Scan(&userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User Not Found"})
			return
		}
		_, err = db.Exec("UPDATE users SET first_name=$1, last_name=$2, email=$3, age=$4 WHERE id=$5", newUser.FirstName, newUser.LastName, newUser.Email, newUser.Age, newUser.ID)
		
		if err !nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}
	} else {
		userID = generateID()
		_, err := db.Exec("INSERT INTO users (id, first_name, last_name, email, age) VALUES ($1, $2, $3, $4, $5)", userID, newUser.FirstName, newUser.LastName, newUser.Email, newUser.Age)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
			return
		}
	}
	
		newUser.ID = userID
		c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "updated_user": newUser})

	}

	/*
	if newUser.ID != "" {
		_, exists := users[newUser.ID]
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User Not Found"})
			return
		}
		users[newUser.ID] = newUser
		c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "updated_user": newUser})
		return
	}

	newUser.ID = generateID()
	users[newUser.ID] = newUser
	c.JSON(http.StatusCreated, newUser)
	*/
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
	/*
	userID := c.Param("id")

	user, exists := users[userID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User Not Found"})
		return
	}

	c.JSON(http.StatusOK, user)
	*/
}

func main() {

	initDB()

	router := gin.Default()
	router.POST("/save", saveUser)
	router.GET("/find/:id", findUser)
	router.Run("localhost:8080")
}
