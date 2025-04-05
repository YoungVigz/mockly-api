package services

import (
	"encoding/json"
	"os"
	"os/exec"

	"github.com/YoungVigz/mockly-api/internal/utils"
)

type SchemaService struct {
}

func NewSchemaService() *SchemaService {
	return &SchemaService{}
}

func (s *SchemaService) GenerateSchema(jsonData []byte) ([]byte, error) {

	var jsonValidator interface{}
	if err := json.Unmarshal(jsonData, &jsonValidator); err != nil {
		return nil, &CustomError{Code: 400, ErrorMessage: "Invalid JSON format"}
	}

	fileName := "./schemas/" + utils.RandomString(10) + ".json"

	if err := os.WriteFile(fileName, jsonData, 0644); err != nil {
		return nil, &CustomError{Code: 400, ErrorMessage: "Could not create a JSON file"}
	}

	cmd := exec.Command("mockly", "generate", "-s", fileName)

	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		return nil, &CustomError{
			Code:         400,
			ErrorMessage: "CLI tool error: " + err.Error() + ": " + string(outputBytes),
		}
	}

	outputFile := "data.json"
	outputData, err := os.ReadFile(outputFile)

	if err != nil {
		return nil, &CustomError{
			Code:         400,
			ErrorMessage: "Could not read: " + err.Error(),
		}
	}

	return outputData, nil
}
