package mocks

import (
	"github.com/YoungVigz/mockly-api/internal/models"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) InsertUser(user models.User) (*models.UserResponse, error) {
	args := m.Called(user)
	return args.Get(0).(*models.UserResponse), args.Error(1)
}

func (m *UserRepositoryMock) FindById(id int) (*models.User, error) {
	args := m.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *UserRepositoryMock) FindByUsername(username string) (*models.User, error) {
	args := m.Called(username)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *UserRepositoryMock) FindByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *UserRepositoryMock) DeleteByID(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *UserRepositoryMock) ChangePassword(id int, password string) error {
	args := m.Called(id, password)
	return args.Error(0)
}
