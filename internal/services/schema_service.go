package services

import (
	"encoding/json"
	"os"
	"os/exec"

	"github.com/YoungVigz/mockly-api/internal/models"
	"github.com/YoungVigz/mockly-api/internal/repository"
	"github.com/YoungVigz/mockly-api/internal/utils"
)

type SchemaService struct {
	repo repository.ISchemaRepository
}

func NewSchemaService(repo repository.ISchemaRepository) *SchemaService {
	return &SchemaService{repo: repo}
}

func (s *SchemaService) GenerateFromSchema(jsonData []byte) ([]byte, error) {

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

func (s *SchemaService) SaveSchema(schemaRequest *models.SchemaRequest, userId int) (*models.Schema, error) {
	var schema models.Schema = models.Schema{
		Title:   schemaRequest.Title,
		Content: schemaRequest.Content,
		UserId:  userId,
	}

	// TODO Check if user does not have another schema with the same title

	insertedSchema, err := s.repo.InsertSchema(schema)

	if err != nil {
		return nil, &CustomError{
			Code:         500,
			ErrorMessage: "Could not create schema",
		}
	}

	return insertedSchema, nil
}

func (s *SchemaService) GetAllUserSchemas(userId int) (*[]models.SchemaResponse, error) {

	schemas, err := s.repo.GetAllUserSchemas(userId)

	if err != nil {
		return nil, &CustomError{
			Code:         500,
			ErrorMessage: "Could not read schemas",
		}
	}

	return schemas, nil
}
