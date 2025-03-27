package repository

import (
	"database/sql"

	"github.com/YoungVigz/mockly-api/internal/database"
	"github.com/YoungVigz/mockly-api/internal/models"
)

type BookRepository interface {
	GetAll() ([]models.Book, error)
}

type DBBookRepository struct {
	db *sql.DB
}

func NewDBBookRepository() (*DBBookRepository, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	return &DBBookRepository{db: db}, nil
}

func (repo *DBBookRepository) GetAll() ([]models.Book, error) {
	rows, err := repo.db.Query("SELECT id, title, author FROM book")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
