package v1_test

import (
    "bytes"
    "encoding/json"
    "errors"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/golang/mock/gomock"
    "github.com/isabellecostawex/ps-tag-onboarding-go/internal/api/handlers/v1"
    "github.com/isabellecostawex/ps-tag-onboarding-go/internal/domain/users"
    "github.com/isabellecostawex/ps-tag-onboarding-go/internal/mocks"
    "github.com/isabellecostawex/ps-tag-onboarding-go/internal/services"
    "github.com/stretchr/testify/assert"
)

func setupRouterPost(service *services.UserManagementService) *gin.Engine {
    router := gin.Default()
    userHandler := &v1.UserHandler{UserService: *service}
    router.POST("/save", userHandler.SaveUserHandler)
    return router
}

func TestSaveUserHandler(t *testing.T) {
    t.Run("Valid User", func(t *testing.T) {
        ctrl := gomock.NewController(t)
        defer ctrl.Finish()
        mockRepo := mocks.NewMockUsersRepository(ctrl)
        service := services.UserManagementService{UserRepo: mockRepo}

        newUser := users.UserData{
            FirstName: "Caio",
            LastName:  "Alves",
            Email:     "caio.alves@gmail.com",
            Age:       30,
        }

        mockRepo.EXPECT().CreateUser(&newUser).Return(1, nil)

        recorder := httptest.NewRecorder()
        router := setupRouterPost(&service)

        jsonBody, _ := json.Marshal(newUser)
        req, _ := http.NewRequest(http.MethodPost, "/save", bytes.NewReader(jsonBody))
        req.Header.Set("Content-Type", "application/json")

        router.ServeHTTP(recorder, req)

        assert.Equal(t, http.StatusOK, recorder.Code)

        var responseBody v1.SaveUserResponse
        err := json.Unmarshal(recorder.Body.Bytes(), &responseBody)
        assert.NoError(t, err)
        assert.Equal(t, 1, responseBody.ID)
        assert.Equal(t, "User saved successfully", responseBody.Message)
    })

    t.Run("Invalid User", func(t *testing.T) {
        ctrl := gomock.NewController(t)
        defer ctrl.Finish()
        mockRepo := mocks.NewMockUsersRepository(ctrl)
        service := services.UserManagementService{UserRepo: mockRepo}

        invalidUser := users.UserData{
            FirstName: "Caio",
            LastName:  "Alves",
            Email:     "caio.alves.com",
            Age:       30,
        }

        recorder := httptest.NewRecorder()
        router := setupRouterPost(&service)

        jsonBody, _ := json.Marshal(invalidUser)
        req, _ := http.NewRequest(http.MethodPost, "/save", bytes.NewReader(jsonBody))
        req.Header.Set("Content-Type", "application/json")

        router.ServeHTTP(recorder, req)

        assert.Equal(t, http.StatusBadRequest, recorder.Code)

        var responseBody v1.SaveUserResponse
        err := json.Unmarshal(recorder.Body.Bytes(), &responseBody)
        assert.NoError(t, err)
        assert.Equal(t, "user did not pass validation: User email must be properly formatted", responseBody.Message)
    })

    t.Run("Repository Error", func(t *testing.T) {
        ctrl := gomock.NewController(t)
        defer ctrl.Finish()
        mockRepo := mocks.NewMockUsersRepository(ctrl)
        service := services.UserManagementService{UserRepo: mockRepo}

        newUser := users.UserData{
            FirstName: "Caio",
            LastName:  "Alves",
            Email:     "caio.alves@gmail.com",
            Age:       30,
        }

        mockRepo.EXPECT().CreateUser(&newUser).Return(0, errors.New("failed to save user"))

        recorder := httptest.NewRecorder()
        router := setupRouterPost(&service)

        jsonBody, _ := json.Marshal(newUser)
        req, _ := http.NewRequest(http.MethodPost, "/save", bytes.NewReader(jsonBody))
        req.Header.Set("Content-Type", "application/json")

        router.ServeHTTP(recorder, req)

        assert.Equal(t, http.StatusInternalServerError, recorder.Code)

        var responseBody v1.SaveUserResponse
        err := json.Unmarshal(recorder.Body.Bytes(), &responseBody)
        assert.NoError(t, err)
        assert.Equal(t, "failed to save user", responseBody.Message)
    })
}