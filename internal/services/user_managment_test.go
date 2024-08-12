package services_test

import (
    "errors"
    "testing"

    "github.com/isabellecostawex/ps-tag-onboarding-go/internal/domain/users"
    "github.com/isabellecostawex/ps-tag-onboarding-go/internal/services"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

func TestSaveUser(t *testing.T) {
    t.Run("Valid User", func(t *testing.T) {
        mockRepo := users.NewMock_UsersRepository(t)
        service := services.UserManagementService{UserRepo: mockRepo}

        newUser := users.UserData{
            FirstName: "Caio",
            LastName:  "Alves",
            Email:     "caio.alves@gmail.com",
            Age:       30,
        }

        mockRepo.On("CreateUser", mock.Anything).Return(1, nil)

        userID, err := service.SaveUser(newUser)

        assert.NoError(t, err)
        assert.Equal(t, 1, userID)
    })

    t.Run("User Underage", func(t *testing.T) {
        mockRepo := new(users.Mock_UsersRepository)
        service := services.UserManagementService{UserRepo: mockRepo}

        newUser := users.UserData{
            FirstName: "Caio",
            LastName:  "Alves",
            Email:     "caio.alves@gmail.com",
            Age:       16,
        }

        userID, err := service.SaveUser(newUser)

        assert.Error(t, err)
        assert.Equal(t, "user did not pass validation: User does not meet minimum age requirement", err.Error())
        assert.Equal(t, 0, userID)
    })

    t.Run("Invalid Email", func(t *testing.T) {
        mockRepo := new(users.Mock_UsersRepository)
        service := services.UserManagementService{UserRepo: mockRepo}

        newUser := users.UserData{
            FirstName: "Caio",
            LastName:  "Alves",
            Email:     "caio.alves.com",
            Age:       30,
        }

        userID, err := service.SaveUser(newUser)

        assert.Error(t, err)
        assert.Equal(t, "user did not pass validation: User email must be properly formatted", err.Error())
        assert.Equal(t, 0, userID)
    })

    t.Run("Missing First Name", func(t *testing.T) {
        mockRepo := new(users.Mock_UsersRepository)
        service := services.UserManagementService{UserRepo: mockRepo}

        newUser := users.UserData{
            LastName:  "Alves",
            Email:     "caio.alves@gmail.com",
            Age:       30,
        }

        userID, err := service.SaveUser(newUser)

        assert.Error(t, err)
        assert.Equal(t, "user did not pass validation: User first/last names required", err.Error())
        assert.Equal(t, 0, userID)
    })

    t.Run("Missing Last Name", func(t *testing.T) {
        mockRepo := new(users.Mock_UsersRepository)
        service := services.UserManagementService{UserRepo: mockRepo}

        newUser := users.UserData{
            FirstName: "Caio",
            Email:     "caio.alves@gmail.com",
            Age:       30,
        }

        userID, err := service.SaveUser(newUser)

        assert.Error(t, err)
        assert.Equal(t, "user did not pass validation: User first/last names required", err.Error())
        assert.Equal(t, 0, userID)
    })

    t.Run("Missing Email", func(t *testing.T) {
        mockRepo := new(users.Mock_UsersRepository)
        service := services.UserManagementService{UserRepo: mockRepo}

        newUser := users.UserData{
            FirstName: "Caio",
            LastName:  "Alves",
            Age:       30,
        }

        userID, err := service.SaveUser(newUser)

        assert.Error(t, err)
        assert.Equal(t, "user did not pass validation: User email required", err.Error())
        assert.Equal(t, 0, userID)
    })

    t.Run("Update Existing User", func(t *testing.T) {
        mockRepo := new(users.Mock_UsersRepository)
        service := services.UserManagementService{UserRepo: mockRepo}

        existingUser := users.UserData{
            ID:        1,
            FirstName: "Caio",
            LastName:  "Alves",
            Email:     "caio.alves@gmail.com",
            Age:       30,
        }

        mockRepo.On("UpdateUser", mock.Anything).Return(nil)

        userID, err := service.SaveUser(existingUser)

        assert.NoError(t, err)
        assert.Equal(t, 1, userID)
        mockRepo.AssertExpectations(t)
    })

    t.Run("Repository CreateUser Error", func(t *testing.T) {
        mockRepo := new(users.Mock_UsersRepository)
        service := services.UserManagementService{UserRepo: mockRepo}

        newUser := users.UserData{
            FirstName: "Caio",
            LastName:  "Alves",
            Email:     "caio.alves@gmail.com",
            Age:       30,
        }

        mockRepo.On("CreateUser", mock.Anything).Return(0, errors.New("failed to save user"))

        userID, err := service.SaveUser(newUser)

        assert.Error(t, err)
        assert.Equal(t, "failed to save user", err.Error())
        assert.Equal(t, 0, userID)
        mockRepo.AssertExpectations(t)
    })
}

func TestRetrieveUser(t *testing.T) {
    t.Run("User Found", func(t *testing.T) {
        mockRepo := users.NewMock_UsersRepository(t)
        service := services.UserManagementService{UserRepo: mockRepo}

        expectedUser := users.UserData{
            ID:        1,
            FirstName: "Caio",
            LastName:  "Alves",
            Email:     "caio.alves@gmail.com",
            Age:       30,
        }

        mockRepo.On("RetrieveUser", "1").Return(expectedUser, nil)

        user, err := service.RetrieveUser("1")

        assert.NoError(t, err)
        assert.Equal(t, expectedUser, user)
    })

    t.Run("User Not Found", func(t *testing.T) {
        mockRepo := users.NewMock_UsersRepository(t)
        service := services.UserManagementService{UserRepo: mockRepo}

        mockRepo.On("RetrieveUser", "1").Return(users.UserData{}, errors.New("user not found"))

        user, err := service.RetrieveUser("1")

        assert.Error(t, err)
        assert.Equal(t, "user not found", err.Error())
        assert.Equal(t, users.UserData{}, user)
    })
}
