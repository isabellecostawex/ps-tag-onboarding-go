package services

import (
	"github.com/isabellecostawex/ps-tag-onboarding-go/internal/models"
	"github.com/isabellecostawex/ps-tag-onboarding-go/pkg/postgresql"
)

func UpdateUser(user *models.UserData) error {
	_, err := postgresql.db.Exec("UPDATE users SET first_name=$1, last_name=$2, email=$3, age=$4 WHERE id=$5", user.FirstName, user.LastName, user.Email, user.Age, user.ID)
	return err
}