package daos

import (
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/config"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/models"
)

// BookDAO persists book data in database
type BookDAO struct{}

// NewBookDAO creates a new BookDAO
func NewBookDAO() *BookDAO {
	return &BookDAO{}
}

func (dao *BookDAO) FindAll() []models.Book {
	var books []models.Book
	config.Config.DB.Find(&books)
	return books
}
