package v1_test

import (
    "net/http"
	"encoding/json"
    "net/http/httptest"
    "testing"
	"errors"

    "github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
    "github.com/isabellecostawex/ps-tag-onboarding-go/internal/api/handlers/v1"
    "github.com/isabellecostawex/ps-tag-onboarding-go/internal/domain/users"
    "github.com/isabellecostawex/ps-tag-onboarding-go/internal/mocks"
    "github.com/isabellecostawex/ps-tag-onboarding-go/internal/services"
    "github.com/stretchr/testify/assert"
)

func setupRouterGet (service *services.UserManagementService) *gin.Engine {
    router := gin.Default()
    userHandler := &v1.UserHandler{UserService: *service}
    router.GET("/find/:id", userHandler.FindUserHandler)
    return router
}

func TestFindUserHandler(t *testing.T) {
    t.Run("Valid User", func(t *testing.T) {
        ctrl := gomock.NewController(t)
        defer ctrl.Finish()
        mockRepo := mocks.NewMockUsersRepository(ctrl)
        service := services.UserManagementService{UserRepo: mockRepo}

        expectedUser := users.UserData{
            ID:        1,
            FirstName: "Caio",
            LastName:  "Alves",
            Email:     "caio.alves@gmail.com",
            Age:       30,
        }

        mockRepo.EXPECT().RetrieveUser("1").Return(expectedUser, nil)

        recorder := httptest.NewRecorder()
        router := setupRouterGet(&service)

        req, _ := http.NewRequest(http.MethodGet, "/find/1", nil)
        router.ServeHTTP(recorder, req)

        assert.Equal(t, http.StatusOK, recorder.Code)
        var responseBody v1.FindUserResponse
        err := json.Unmarshal(recorder.Body.Bytes(), &responseBody)
        assert.NoError(t, err)
        assert.Equal(t, expectedUser.ID, responseBody.ID)
        assert.Equal(t, expectedUser.FirstName, responseBody.FirstName)
        assert.Equal(t, expectedUser.LastName, responseBody.LastName)
        assert.Equal(t, expectedUser.Email, responseBody.Email)
        assert.Equal(t, expectedUser.Age, responseBody.Age)
    })

    t.Run("Invalid User", func(t *testing.T) {
        ctrl := gomock.NewController(t)
        defer ctrl.Finish()
        mockRepo := mocks.NewMockUsersRepository(ctrl)
        service := services.UserManagementService{UserRepo: mockRepo}

        mockRepo.EXPECT().RetrieveUser("1").Return(users.UserData{}, errors.New("user not found"))

        recorder := httptest.NewRecorder()
        router := setupRouterGet(&service)

        req, _ := http.NewRequest(http.MethodGet, "/find/1", nil)
        router.ServeHTTP(recorder, req)

        assert.Equal(t, http.StatusNotFound, recorder.Code)
        var responseBody map[string]string
        err := json.Unmarshal(recorder.Body.Bytes(), &responseBody)
        assert.NoError(t, err)
        assert.Equal(t, "user not found", responseBody["error"])
    })

    t.Run("Repository Error", func(t *testing.T) {
        ctrl := gomock.NewController(t)
        defer ctrl.Finish()
        mockRepo := mocks.NewMockUsersRepository(ctrl)
        service := services.UserManagementService{UserRepo: mockRepo}

        mockRepo.EXPECT().RetrieveUser("1").Return(users.UserData{}, errors.New("repository error"))

        recorder := httptest.NewRecorder()
        router := setupRouterGet(&service)

        req, _ := http.NewRequest(http.MethodGet, "/find/1", nil)
        router.ServeHTTP(recorder, req)

        assert.Equal(t, http.StatusInternalServerError, recorder.Code)
        var responseBody map[string]string
        err := json.Unmarshal(recorder.Body.Bytes(), &responseBody)
        assert.NoError(t, err)
        assert.Equal(t, "repository error", responseBody["error"])
    })
}
