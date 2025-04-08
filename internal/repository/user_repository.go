package repository

import (
	"database/sql"
	"fmt"

	"github.com/YoungVigz/mockly-api/internal/database"
	"github.com/YoungVigz/mockly-api/internal/models"
)

type IUserRepository interface {
	InsertUser(models.User) (*models.UserResponse, error)
	FindById(int) (*models.User, error)
	FindByUsername(string) (*models.User, error)
	FindByEmail(string) (*models.User, error)
	DeleteByID(int) error
	ChangePassword(int, string) error
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

func (repo *UserRepository) FindById(user_id int) (*models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE id = $1 LIMIT 1"
	row := repo.db.QueryRow(query, user_id)

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil

}

func (repo *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE username = $1 LIMIT 1"
	row := repo.db.QueryRow(query, username)

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE email = $1 LIMIT 1"
	row := repo.db.QueryRow(query, email)

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
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

func (repo *UserRepository) DeleteByID(userId int) error {

	query := `DELETE FROM users WHERE id = $1`
	result, err := repo.db.Exec(query, userId)
	if err != nil {
		return fmt.Errorf("DeleteByID error: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeleteByID error: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("DeleteByID error: no user found with id %d", userId)
	}

	return nil
}

func (repo *UserRepository) ChangePassword(userId int, hashPassword string) error {

	query := `UPDATE users SET password = $1 WHERE id = $2`
	result, err := repo.db.Exec(query, hashPassword, userId)
	if err != nil {
		return fmt.Errorf("ChangePassword error: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ChangePassword error: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("ChangePassword error: no user found with id %d", userId)
	}

	return nil
}
