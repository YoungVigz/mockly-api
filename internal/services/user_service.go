package services

import (
	"fmt"

	"github.com/YoungVigz/mockly-api/internal/models"
	"github.com/YoungVigz/mockly-api/internal/repository"
	"github.com/YoungVigz/mockly-api/internal/utils"
)

type CustomError struct {
	Code         int
	ErrorMessage string
}

func (e *CustomError) Error() string {
	return e.ErrorMessage
}

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(userData *models.UserAuthRequest) (*models.UserResponse, error) {
	var user models.User = models.User{Username: userData.Username, Email: userData.Email}

	existingUser, err := s.repo.FindByUsername(user.Username)

	if existingUser != nil {
		return nil, &CustomError{Code: 409, ErrorMessage: "Username already in use"}
	} else if err != nil {
		return nil, &CustomError{Code: 500, ErrorMessage: "Internal Server Error"}
	}

	existingUser, err = s.repo.FindByEmail(user.Email)

	if existingUser != nil {
		return nil, &CustomError{Code: 409, ErrorMessage: "Email already in use"}
	} else if err != nil {
		return nil, &CustomError{Code: 500, ErrorMessage: "Internal Server Error"}
	}

	hash, err := utils.HashPassword(userData.Password)

	if err != nil {
		return nil, &CustomError{Code: 500, ErrorMessage: "Internal Server Error"}
	}

	user.Password = hash

	createdUser, err := s.repo.InsertUser(user)

	if err != nil {
		return nil, &CustomError{Code: 500, ErrorMessage: "Internal Server Error"}
	}

	return createdUser, nil
}

func (s *UserService) Login(userData *models.UserLoginRequest) (string, error) {

	user, err := s.repo.FindByEmail(userData.Email)

	if err != nil {
		return "", &CustomError{Code: 500, ErrorMessage: "Internal Server Error"}
	}

	if user == nil {
		return "", &CustomError{Code: 401, ErrorMessage: "Invalid email, or password"}
	}

	isPasswordValid := utils.CheckPasswordHash(userData.Password, user.Password)

	if !isPasswordValid {
		return "", &CustomError{Code: 401, ErrorMessage: "Invalid email, or password"}
	}

	token, err := utils.CreateJWTToken(user)

	fmt.Println(err)

	if err != nil {
		return "", &CustomError{Code: 500, ErrorMessage: "Internal Server Error"}
	}

	return token, nil
}

func (s *UserService) GetUserById(userId int) (*models.UserResponse, error) {
	user, err := s.repo.FindById(userId)

	if err != nil {
		return nil, &CustomError{Code: 500, ErrorMessage: "Internal Server Error"}
	}

	return &models.UserResponse{Id: user.Id, Username: user.Username, Email: user.Email}, nil
}

func (s *UserService) DeleteUserById(userId int, userRequest *models.UserModifyRequest) error {

	user, err := s.repo.FindById(userId)

	if err != nil {
		return &CustomError{Code: 404, ErrorMessage: "Could not find user"}
	}

	isPasswordValid := utils.CheckPasswordHash(userRequest.Password, user.Password)

	if !isPasswordValid {
		return &CustomError{Code: 401, ErrorMessage: "Unauthorized action"}
	}

	err = s.repo.DeleteByID(userId)

	if err != nil {
		return &CustomError{Code: 500, ErrorMessage: "Internal Server Error"}
	}

	return nil
}

func (s *UserService) ChangePassword(userId int, userRequest *models.UserChangePassword) error {

	user, err := s.repo.FindById(userId)

	if err != nil {
		return &CustomError{Code: 404, ErrorMessage: "Could not find user"}
	}

	isPasswordValid := utils.CheckPasswordHash(userRequest.Password, user.Password)

	if !isPasswordValid {
		return &CustomError{Code: 401, ErrorMessage: "Unauthorized action"}
	}

	hashPassword, err := utils.HashPassword(userRequest.NewPassword)

	if err != nil {
		return &CustomError{Code: 500, ErrorMessage: "Internal Server Error"}
	}

	err = s.repo.ChangePassword(userId, hashPassword)

	if err != nil {
		return &CustomError{Code: 500, ErrorMessage: "Internal Server Error"}
	}

	return nil
}
