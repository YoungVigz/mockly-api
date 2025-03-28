package handlers

import (
	"net/http"

	"github.com/YoungVigz/mockly-api/internal/models"
	"github.com/YoungVigz/mockly-api/internal/validators"
	"github.com/gin-gonic/gin"
)

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
	}

}

func LoginUser(c *gin.Context) {

}
