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

func RegisterUser(c *gin.Context) {

	userCreateRequest := &models.UserCreateRequest{}

	if c.Bind(&userCreateRequest) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Could not read provided values, ensure that your body is correct",
		})

		return
	}

	validatorMasseges, err := validators.CreateUserValidator(userCreateRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"errors": validatorMasseges,
		})

		return
	}

	createduser, err := userService.CreateUser(userCreateRequest)

	if err != nil {
		customError := err.(*services.CustomError)

		c.JSON(customError.Code, gin.H{
			"code":  customError.Code,
			"error": customError.ErrorMessage,
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code": http.StatusCreated,
		"data": createduser,
	})

}

func LoginUser(c *gin.Context) {

}
