package services

import (
	"github.com/isabellecostawex/ps-tag-onboarding-go/pkg/postgresql"
	"github.com/isabellecostawex/ps-tag-onboarding-go/internal/models"
)

func CreateUser(user *models.UserData) (int, error){
	var lastinsertID int
	err := postgresql.db.QueryRow("INSERT INTO users (first_name, last_name, email, age) VALUES ($1, $2, $3, $4) RETURNING id", user.FirstName, user.LastName, user.Email, user.Age).Scan(&lastinsertID)
	if err != nil {
		return 0, err
	}
	return lastinsertID, nil
}
