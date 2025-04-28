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
	GetUserSchemaByTitle(string, int) (*models.Schema, error)
	DeleteSchema(string, int) error
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

func (repo *SchemaRepository) GetUserSchemaByTitle(title string, userId int) (*models.Schema, error) {
	var schema models.Schema

	query := `SELECT id, title, content, user_id FROM schemas WHERE user_id = $1 AND title = $2 LIMIT 1`

	err := repo.db.QueryRow(query, userId, title).Scan(&schema.Id, &schema.Title, &schema.Content, &schema.UserId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &schema, nil
}

func (repo *SchemaRepository) DeleteSchema(title string, userId int) error {
	query := `DELETE FROM schemas WHERE user_id = $1 AND title = $2`

	result, err := repo.db.Exec(query, userId, title)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
