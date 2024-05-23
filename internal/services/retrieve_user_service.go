package services

import (
	"database/sql"
	"errors"

	"github.com/isabellecostawex/ps-tag-onboarding-go/internal/models"
	"github.com/isabellecostawex/ps-tag-onboarding-go/pkg/postgresql"
)

func RetrieveUser(userID string)(models.UserData, error) {
	var user models.UserData
	err := postgresql.DB.QueryRow("SELECT id, first_name, last_name, email, age FROM users WHERE ID=$1", userID).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("User not found")
		}
		return user, err
	}
	return user, nil
}
