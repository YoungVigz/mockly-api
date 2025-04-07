package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/YoungVigz/mockly-api/internal/models"
	"github.com/YoungVigz/mockly-api/internal/repository"
	"github.com/YoungVigz/mockly-api/internal/services"
	"github.com/YoungVigz/mockly-api/internal/utils"
	"github.com/YoungVigz/mockly-api/internal/validators"
	"github.com/gin-gonic/gin"
)

var schemaService services.SchemaService

func init() {
	schemaRepository, _ := repository.NewSchemaRepository()
	schemaService = *services.NewSchemaService(schemaRepository)
}

func GenerateFromSchema(c *gin.Context) {

	jsonData, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Invalid Json format",
		})
	}

	c.Request.Body.Close()

	data, err := schemaService.GenerateFromSchema(jsonData)

	if err != nil {
		customError := err.(*services.CustomError)

		c.JSON(customError.Code, gin.H{
			"code":  customError.Code,
			"error": customError.ErrorMessage,
		})

		return
	}

	resp := map[string]json.RawMessage{
		"data": data,
	}

	c.JSON(http.StatusOK, resp)
}

func SaveSchema(c *gin.Context) {
	schemaCreateRequest := &models.SchemaRequest{}

	if c.Bind(&schemaCreateRequest) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Could not read provided values, ensure that your body is correct",
		})

		return
	}

	validatorMasseges, err := validators.SchemaValidator(schemaCreateRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"errors": validatorMasseges,
		})

		return
	}

	userId, exist := c.Get("user_id")

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":   http.StatusUnauthorized,
			"errors": "Invalid token encoding, please log in again",
		})

		return
	}

	userIdInt, err := utils.ConvertUserIdToInt(userId)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":   http.StatusUnauthorized,
			"errors": "Invalid token encoding, please log in again",
		})

		return
	}

	schema, err := schemaService.SaveSchema(schemaCreateRequest, userIdInt)

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
		"data": schema,
	})
}

func GetAllUserSchemas(c *gin.Context) {

	userId, exist := c.Get("user_id")

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":   http.StatusUnauthorized,
			"errors": "Invalid token encoding, please log in again",
		})

		return
	}

	userIdInt, err := utils.ConvertUserIdToInt(userId)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":   http.StatusUnauthorized,
			"errors": "Invalid token encoding, please log in again",
		})

		return
	}

	schemas, err := schemaService.GetAllUserSchemas(userIdInt)

	if err != nil {
		customError := err.(*services.CustomError)

		c.JSON(customError.Code, gin.H{
			"code":  customError.Code,
			"error": customError.ErrorMessage,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": schemas,
	})
}
