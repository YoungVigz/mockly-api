package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/YoungVigz/mockly-api/internal/handlers"
	"github.com/YoungVigz/mockly-api/internal/mocks"
	"github.com/YoungVigz/mockly-api/internal/models"
	"github.com/YoungVigz/mockly-api/internal/routes"
	"github.com/YoungVigz/mockly-api/internal/services"
	"github.com/YoungVigz/mockly-api/internal/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUser_WithMock(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryMock)
	mockService := *services.NewUserService(mockRepo)
	handlers.SetUserService(mockService)

	userRequest := models.UserAuthRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "Test1234!",
	}

	body, _ := json.Marshal(userRequest)

	mockRepo.On("FindByUsername", "testuser").Return((*models.User)(nil), nil)
	mockRepo.On("FindByEmail", "test@example.com").Return((*models.User)(nil), nil)
	mockRepo.On("InsertUser", mock.AnythingOfType("models.User")).Return(&models.UserResponse{
		Id:       1,
		Username: "testuser",
		Email:    "test@example.com",
	}, nil)

	req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := routes.SetupTestRoutes()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestLogin_Success(t *testing.T) {

	pass, _ := utils.HashPassword("Test1234!")
	mockRepo := new(mocks.UserRepositoryMock)
	service := services.NewUserService(mockRepo)

	mockUser := &models.User{
		Id:       1,
		Email:    "user@example.com",
		Password: pass,
	}

	mockRepo.On("FindByEmail", "user@example.com").Return(mockUser, nil)

	req := &models.UserLoginRequest{
		Email:    "user@example.com",
		Password: "Test1234!",
	}

	token, err := service.Login(req)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	mockRepo.AssertExpectations(t)
}
