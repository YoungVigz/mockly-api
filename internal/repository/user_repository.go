package repository

import (
	"database/sql"
	"fmt"

	"github.com/YoungVigz/mockly-api/internal/database"
	"github.com/YoungVigz/mockly-api/internal/models"
)

type IUserRepository interface {
	InsertUser(models.User) (*models.UserResponse, error)
	FindByUsername(string) (*models.UserResponse, error)
	FindByEmail(string) (*models.UserResponse, error)
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

func (repo *UserRepository) FindByUsername(username string) (*models.UserResponse, error) {
	var user models.UserResponse

	query := "SELECT id, username, email FROM users WHERE username = $1 LIMIT 1"
	row := repo.db.QueryRow(query, username)

	err := row.Scan(&user.Id, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) FindByEmail(email string) (*models.UserResponse, error) {
	var user models.UserResponse

	query := "SELECT id, username, email FROM users WHERE email = $1 LIMIT 1"
	row := repo.db.QueryRow(query, email)

	err := row.Scan(&user.Id, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) InsertUser(user models.User) (*models.UserResponse, error) {
	var newUser models.UserResponse

	query := `INSERT INTO users (username, email, password) 
	          VALUES ($1, $2, $3) RETURNING id, username, email`

	err := repo.db.QueryRow(query, user.Username, user.Email, user.Password).
		Scan(&newUser.Id, &newUser.Username, &newUser.Email)

	if err != nil {
		return nil, fmt.Errorf("InsertUser error: %w", err)
	}

	return &newUser, nil
}
