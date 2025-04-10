package services

import (
	"encoding/json"
	"os"
	"os/exec"

	"github.com/YoungVigz/mockly-api/internal/models"
	"github.com/YoungVigz/mockly-api/internal/repository"
	"github.com/YoungVigz/mockly-api/internal/utils"
	"github.com/YoungVigz/mockly-api/internal/websockets"
)

type SchemaService struct {
	repo repository.ISchemaRepository
}

func NewSchemaService(repo repository.ISchemaRepository) *SchemaService {
	return &SchemaService{repo: repo}
}

func (s *SchemaService) GenerateFromSchema(jsonData []byte, userId int) ([]byte, error) {

	var jsonValidator interface{}
	if err := json.Unmarshal(jsonData, &jsonValidator); err != nil {
		websockets.SendToUser(userId, []byte("❌ Invalid JSON payload received. Generation canceled."))

		return nil, &CustomError{Code: 400, ErrorMessage: "Invalid JSON format"}
	}

	fileName := utils.RandomString(10)
	schemaFilePath := "./schemas/" + fileName + ".json"
	outputFilePath := fileName + ".json"

	if err := os.WriteFile(schemaFilePath, jsonData, 0644); err != nil {
		websockets.SendToUser(userId, []byte("❌ Could not create a JSON file. Generation canceled."))
		return nil, &CustomError{Code: 400, ErrorMessage: "Could not create a JSON file"}
	}

	cmd := exec.Command("mockly", "generate", "-s", schemaFilePath, "-o", fileName)

	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		websockets.SendToUser(userId, []byte("❌ Your schema appears to be invalid. Please check and try again."))
		return nil, &CustomError{
			Code:         400,
			ErrorMessage: "CLI tool error: " + err.Error() + ": " + string(outputBytes),
		}
	}

	outputData, err := os.ReadFile(outputFilePath)

	websockets.SendToUser(userId, []byte("✅ Data generated, now retriving data"))

	if err != nil {
		websockets.SendToUser(userId, []byte("❌ Could not read data from generation. Please try again or contact the administrator."))
		return nil, &CustomError{
			Code:         500,
			ErrorMessage: "Could not read: " + err.Error(),
		}
	}

	defer os.Remove(schemaFilePath)
	defer os.Remove(outputFilePath)

	return outputData, nil
}

func (s *SchemaService) SaveSchema(schemaRequest *models.SchemaRequest, userId int) (*models.Schema, error) {
	var schema models.Schema = models.Schema{
		Title:   schemaRequest.Title,
		Content: schemaRequest.Content,
		UserId:  userId,
	}

	existing, err := s.repo.GetUserSchemaByTitle(schemaRequest.Title, userId)

	if err != nil {
		return nil, &CustomError{
			Code:         500,
			ErrorMessage: "Could not check for duplicate title",
		}
	}
	if existing != nil {
		return nil, &CustomError{
			Code:         400,
			ErrorMessage: "You already have a schema with this title",
		}
	}

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

func (s *SchemaService) GetUserSchemaByTitle(title string, userId int) (*models.SchemaResponse, error) {

	schema, err := s.repo.GetUserSchemaByTitle(title, userId)

	if err != nil {
		return nil, &CustomError{
			Code:         500,
			ErrorMessage: "Internal Server Error",
		}
	}

	if schema == nil {
		return nil, &CustomError{
			Code:         404,
			ErrorMessage: "Schema not found",
		}
	}

	var schemaResponse = &models.SchemaResponse{
		Id:      schema.Id,
		Title:   schema.Title,
		Content: schema.Content,
	}

	return schemaResponse, nil
}
