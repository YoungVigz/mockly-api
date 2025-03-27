package services

import (
	"github.com/YoungVigz/mockly-api/internal/models"
	"github.com/YoungVigz/mockly-api/internal/repository"
)

type BookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetBooks() ([]models.Book, error) {
	return s.repo.GetAll()
}
