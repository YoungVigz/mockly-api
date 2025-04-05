package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/YoungVigz/mockly-api/internal/services"
	"github.com/gin-gonic/gin"
)

var schemaService services.SchemaService

func init() {
	schemaService = *services.NewSchemaService()
}

func GenerateSchema(c *gin.Context) {

	jsonData, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Invalid Json format",
		})
	}

	c.Request.Body.Close()

	data, err := schemaService.GenerateSchema(jsonData)

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
