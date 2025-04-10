package validators

import (
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"regexp"

	"github.com/YoungVigz/mockly-api/internal/models"
	"github.com/YoungVigz/mockly-api/internal/services"
	"github.com/YoungVigz/mockly-api/internal/utils"
)

func SchemaValidator(schemaRequest *models.SchemaRequest) error {

	var jsonValidator interface{}
	if err := json.Unmarshal(schemaRequest.Content, &jsonValidator); err != nil {
		return &services.CustomError{Code: 400, ErrorMessage: "Invalid JSON format"}
	}

	fileName := utils.RandomString(10)
	schemaFilePath := "./schemas/" + fileName + ".json"

	if err := os.WriteFile(schemaFilePath, schemaRequest.Content, 0644); err != nil {
		return &services.CustomError{Code: 400, ErrorMessage: "Could not create a JSON file"}
	}

	cmd := exec.Command("mockly", "test", "-s", schemaFilePath)

	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		return &services.CustomError{
			Code:         400,
			ErrorMessage: "CLI tool error: " + string(outputBytes),
		}
	}

	defer os.Remove(schemaFilePath)

	return nil
}

var titleRegex = regexp.MustCompile(`^[a-zA-Z0-9_]{2,20}$`)

func TitleValidator(title string) ([]string, error) {
	var validationErrors []string

	if title == "" {
		validationErrors = append(validationErrors, "Title is required")
	} else if !titleRegex.MatchString(title) {
		validationErrors = append(validationErrors, "Title must be 2-20 characters long and can only contain letters, numbers, and underscores")
	}

	if len(validationErrors) > 0 {
		return validationErrors, errors.New("validation failed")
	}

	return nil, nil
}
