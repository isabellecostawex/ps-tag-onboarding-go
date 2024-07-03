
package handlers_test

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/isabellecostawex/ps-tag-onboarding-go/internal/api/handlers"
    "github.com/isabellecostawex/ps-tag-onboarding-go/internal/domain/users"
    "github.com/isabellecostawex/ps-tag-onboarding-go/internal/mocks"
    "github.com/isabellecostawex/ps-tag-onboarding-go/internal/services"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)
func setupRouter (service *services.UserManagementService) *gin.Engine {
    router := gin.Default()
    router.POST("/save", handlers.UserHandler(service))
    return router
}

func TestSaveUserHandler(t *testing.T) {
    t.Run("Valid User", func(t *testing.T) {
        mockRepo := new(mocks.UsersRepository)
        service := services.UserManagementService{UserRepo: mockRepo}

        newUser := users.UserData{
            FirstName: "Caio",
            LastName:  "Alves",
            Email:     "caio.alves@gmail.com",
            Age:       30,
        }

        mockRepo.On("CreateUser", mock.Anything).Return(1, nil)

        recorder := httptest.NewRecorder()
        ginContext, _ := gin.CreateTestContext(recorder)
        jsonBody, _ := json.Marshal(newUser)
        ginContext.Request, _ = http.NewRequest(http.MethodPost, "/save", bytes.NewReader(jsonBody))

        UserHandler := UserHandler{UserService: service}
        router := RegisterRoutes(&UserHandler)
        router.ServeHTTP(recorder, ginContext.Request)


        assert.Equal(t, http.StatusOK, recorder.Code)

        var responseBody SaveUserResponse
        err := json.Unmarshal(recorder.Body.Bytes(), &responseBody)
        assert.NoError(t, err)
        assert.Equal(t, 1, responseBody.ID)
        assert.Equal(t, "User saved successfully", responseBody.Message)

        mockRepo.AssertExpectations(t)
    })
}
