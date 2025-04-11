package handlers

import (
	"net/http"

	"github.com/YoungVigz/mockly-api/internal/models"
	"github.com/YoungVigz/mockly-api/internal/repository"
	"github.com/YoungVigz/mockly-api/internal/services"
	"github.com/YoungVigz/mockly-api/internal/validators"
	"github.com/gin-gonic/gin"
)

var userService services.UserService

func init() {
	repo, _ := repository.NewUserRepository()
	userService = *services.NewUserService(repo)
}

// RegisterUser godoc
// @Summary      Create an account
// @Description  Creates a new user account in the system
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user  body      models.UserAuthRequest  true  "User registration data"
// @Success      201   {object}  models.UserResponse
// @Failure      400   {object}  models.ErrorResponse "Invalid request body or validation errors"
// @Failure      409   {object}  models.ErrorResponse "Username or email already in use"
// @Router       /auth/register [post]
func RegisterUser(c *gin.Context) {

	userCreateRequest := &models.UserAuthRequest{}

	if c.Bind(&userCreateRequest) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not read provided values, ensure that your body is correct",
		})

		return
	}

	validatorMasseges, err := validators.CreateUserValidator(userCreateRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": validatorMasseges,
		})

		return
	}

	createduser, err := userService.CreateUser(userCreateRequest)

	if err != nil {
		customError := err.(*services.CustomError)

		c.JSON(customError.Code, gin.H{
			"error": customError.ErrorMessage,
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": createduser,
	})

}

// LoginUser godoc
// @Summary      Log in to get auth token
// @Description  Authenticates user and returns a JWT token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      models.UserLoginRequest  true  "Login credentials"
// @Success      201  {object}  models.Token  "JWT token"
// @Failure      400  {object}  models.ErrorResponse "Invalid request body"
// @Failure      401  {object}  models.ErrorResponse "Invalid credentials"
// @Router       /auth/login [post]
func LoginUser(c *gin.Context) {
	userLoginRequest := &models.UserLoginRequest{}

	if c.Bind(&userLoginRequest) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not read provided values, ensure that your body is correct",
		})

		return
	}

	validatorMasseges, err := validators.LoginValidator(userLoginRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": validatorMasseges,
		})

		return
	}

	token, err := userService.Login(userLoginRequest)

	if err != nil {
		customError := err.(*services.CustomError)

		c.JSON(customError.Code, gin.H{
			"error": customError.ErrorMessage,
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"token": token,
	})

}
