package repository

import (
	"database/sql"

	"github.com/YoungVigz/mockly-api/internal/database"
	"github.com/YoungVigz/mockly-api/internal/models"
)

type IUserRepository interface {
	CreateUser(models.UserCreateRequest) (*models.UserResponse, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() (*UserRepository, error) {
	db, err := database.GetDB()

	if err != nil {
		return nil, err
	}

	return &UserRepository{db: db}, nil
}

func (repo *UserRepository) CreateUser(user models.UserCreateRequest) (*models.UserResponse, error) {

	row, err := repo.db.Query(`INSERT`)

	row.Scan()

	if err != nil {
		return nil, err
	}

	return &models.UserResponse{}, nil

}
