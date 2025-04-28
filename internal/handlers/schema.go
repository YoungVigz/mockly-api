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
	"github.com/YoungVigz/mockly-api/internal/websockets"
	"github.com/gin-gonic/gin"
)

var schemaService services.SchemaService

func init() {
	schemaRepository, _ := repository.NewSchemaRepository()
	schemaService = *services.NewSchemaService(schemaRepository)
}

type SchemaWithJSONResponse struct {
	Id      int               `json:"schema_id"`
	Title   string            `json:"title"`
	Content map[string]string `json:"content"`
}

type SchemaWithJSONRequest struct {
	Title   string
	Content map[string]string
}

// @Summary Generate data from schema
// @Description Accepts a JSON schema and generates data using the CLI tool. Returns the generated data or an error if invalid. To learn more about schemas visit https://github.com/YoungVigz/mockly-cli/blob/main/README.md
// @Tags Schema
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param schema body map[string]string true "JSON Schema for generation"
// @Success 200 {object} map[string]string "Generated data"
// @Failure 400 {object} models.ErrorResponse "Invalid JSON format, or invalid schema syntax"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /schema/generate [post]
func GenerateFromSchema(c *gin.Context) {

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

	websockets.SendToUser(userIdInt, []byte("✅ Authenticated, parsing provided data"))

	jsonData, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Invalid Json format",
		})

		websockets.SendToUser(userIdInt, []byte("❌ Invalid JSON payload received. Generation canceled."))

		return
	}

	c.Request.Body.Close()

	data, err := schemaService.GenerateFromSchema(jsonData, userIdInt)

	if err != nil {
		customError := err.(*services.CustomError)

		c.JSON(customError.Code, gin.H{
			"code":  customError.Code,
			"error": customError.ErrorMessage,
		})

		return
	}

	websockets.SendToUser(userIdInt, []byte("🎉 Operation succesful, sending data"))

	resp := map[string]json.RawMessage{
		"data": data,
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary Save a new schema
// @Description Saves a new schema provided by the user. The schema contains a title and content in JSON format. To learn more about schemas visit https://github.com/YoungVigz/mockly-cli/blob/main/README.md
// @Tags Schema
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param schema body SchemaWithJSONRequest true "Schema data to save"
// @Success 201 {object} SchemaWithJSONResponse "Schema created"
// @Failure 400 {object} models.ErrorResponse "Invalid schema data or duplicate title"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /schema [post]
func SaveSchema(c *gin.Context) {
	schemaCreateRequest := &models.SchemaRequest{}

	if c.Bind(&schemaCreateRequest) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Could not read provided values, ensure that your body is correct",
		})

		return
	}

	err := validators.SchemaValidator(schemaCreateRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	validationMessages, err := validators.TitleValidator(schemaCreateRequest.Title)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": validationMessages,
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

// @Summary Get all schemas for the authenticated user
// @Description Retrieves all schemas created by the authenticated user.
// @Tags Schema
// @Produce json
// @Security BearerAuth
// @Success 200 {array} SchemaWithJSONResponse "List of schemas"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /schema [get]
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

type TitleUri struct {
	Title string `uri:"title" binding:"required"`
}

// @Summary Get user's schema by title
// @Description Retrieves a schema based on the title for the authenticated user.
// @Tags Schema
// @Produce json
// @Security BearerAuth
// @Param title path string true "Schema title"
// @Success 200 {object} SchemaWithJSONResponse "Schema data"
// @Failure 400 {object} models.ErrorResponse "Invalid title was provided"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "Schema not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /schema/{title} [get]
func GetUserSchemaByTitle(c *gin.Context) {

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

	var uri TitleUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	schema, err := schemaService.GetUserSchemaByTitle(uri.Title, userIdInt)

	if err != nil {
		customError := err.(*services.CustomError)

		c.JSON(customError.Code, gin.H{
			"error": customError.ErrorMessage,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": schema,
	})

}

func DeleteUserSchema(c *gin.Context) {
	userId, exist := c.Get("user_id")

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errors": "Invalid token encoding, please log in again",
		})

		return
	}

	userIdInt, err := utils.ConvertUserIdToInt(userId)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errors": "Invalid token encoding, please log in again",
		})

		return
	}

	_, err = userService.GetUserById(userIdInt)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"errors": "The provided token corresponds to a user that no longer exists.",
		})

		return
	}

	var uri TitleUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = schemaService.DeleteUserSchema(uri.Title, userIdInt)

	if err != nil {
		customError := err.(*services.CustomError)

		c.JSON(customError.Code, gin.H{
			"error": customError.ErrorMessage,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Schema deleted succesfully",
	})

}
