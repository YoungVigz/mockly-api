package repository

import (
	"database/sql"
	"fmt"

	"github.com/YoungVigz/mockly-api/internal/database"
	"github.com/YoungVigz/mockly-api/internal/models"
)

type ISchemaRepository interface {
	InsertSchema(models.Schema) (*models.Schema, error)
	GetAllUserSchemas(int) (*[]models.SchemaResponse, error)
}

type SchemaRepository struct {
	db *sql.DB
}

func NewSchemaRepository() (*SchemaRepository, error) {
	db, err := database.GetDB()

	if err != nil {
		return nil, err
	}

	return &SchemaRepository{db: db}, nil
}

func (repo *SchemaRepository) InsertSchema(schema models.Schema) (*models.Schema, error) {

	var newSchema models.Schema

	query := `INSERT INTO schemas (title, user_id, content) 
	          VALUES ($1, $2, $3) RETURNING id, title, user_id, content`

	err := repo.db.QueryRow(query, schema.Title, schema.UserId, schema.Content).
		Scan(&newSchema.Id, &newSchema.Title, &newSchema.UserId, &newSchema.Content)

	if err != nil {
		return nil, fmt.Errorf("InsertSchema error: %w", err)
	}

	return &newSchema, nil
}

func (repo *SchemaRepository) GetAllUserSchemas(userId int) (*[]models.SchemaResponse, error) {

	var schemas []models.SchemaResponse

	query := "SELECT id, title, content FROM schemas WHERE user_id = $1"
	rows, err := repo.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var schema models.SchemaResponse
		err := rows.Scan(&schema.Id, &schema.Title, &schema.Content)
		if err != nil {
			return nil, err
		}
		schemas = append(schemas, schema)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &schemas, nil
}
