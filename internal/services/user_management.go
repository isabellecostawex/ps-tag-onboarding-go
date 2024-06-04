package services

import (
	"errors"
	"strings"

	"github.com/isabellecostawex/ps-tag-onboarding-go/internal/domain/users"
)

type UserManagementService struct {
	UserRepo users.UsersRepository
}

func (ums *UserManagementService) SaveUser(newUser users.UserData) (int, error) {
	if newUser.Age < 18 {
		return 0, errors.New("user did not pass validation: User does not meet minimum age requirement")
	}

	if newUser.FirstName == "" || newUser.LastName == "" {
		return 0, errors.New("user did not pass validation: User first/last names required")
	}

	if newUser.Email == "" {
		return 0, errors.New("user did not pass validation: User email required")
	}

	if !strings.Contains(newUser.Email, "@") {
		return 0, errors.New("user did not pass validation: User email must be properly formatted")
	}

	if newUser.ID != 0 {
		err := ums.UserRepo.UpdateUser(&newUser)
		if err != nil {
			return 0, errors.New("failed to update user")
		}
		return newUser.ID, nil
	} else {
		userID, err := ums.UserRepo.CreateUser(&newUser)
		if err != nil {
			return 0, errors.New("failed to save user")
		}
		return userID, nil
	}
}

func (ums *UserManagementService) RetrieveUser(userID string) (users.UserData, error) {
	return ums.UserRepo.RetrieveUser(userID)
}
