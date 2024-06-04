package postgres

import (
	"database/sql"
	"errors"

	"github.com/isabellecostawex/ps-tag-onboarding-go/internal/domain/users"
	"github.com/isabellecostawex/ps-tag-onboarding-go/pkg/postgresql"
)
type UsersRepository struct {}

func (ur *UsersRepository) CreateUser(user *users.UserData) (int, error) {
	var lastinsertID int
	err := postgresql.DB.QueryRow("INSERT INTO users (first_name, last_name, email, age) VALUES ($1, $2, $3, $4) RETURNING id", user.FirstName, user.LastName, user.Email, user.Age).Scan(&lastinsertID)
	if err != nil {
		return 0, err
	}
	return lastinsertID, nil
}

func (ur *UsersRepository) RetrieveUser(userID string) (users.UserData, error) {
	var user users.UserData
	err := postgresql.DB.QueryRow("SELECT id, first_name, last_name, email, age FROM users WHERE ID=$1", userID).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("user not found")
		}
		return user, err
	}
	return user, nil
}

func (ur *UsersRepository) UpdateUser(user *users.UserData) error {
	_, err := postgresql.DB.Exec("UPDATE users SET first_name=$1, last_name=$2, email=$3, age=$4 WHERE id=$5", user.FirstName, user.LastName, user.Email, user.Age, user.ID)
	return err
}
