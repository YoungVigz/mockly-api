package handlers

import (
	"net/http"

	"github.com/YoungVigz/mockly-api/internal/repository"
	"github.com/YoungVigz/mockly-api/internal/services"
	"github.com/gin-gonic/gin"
)

var bookService services.BookService

func init() {
	repo, _ := repository.NewDBBookRepository()
	bookService = *services.NewBookService(repo)
}

func GetAllBooks(c *gin.Context) {
	books, err := bookService.GetBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}
