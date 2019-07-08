package services

import (
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBookService(t *testing.T) {
	dao := newMockBookDAO()
	s := NewBookService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestBookService_FindAll(t *testing.T) {
	s := NewBookService(newMockBookDAO())
	books, err := s.FindAll()
	if assert.Nil(t, err) && assert.NotEmpty(t, books) {
		assert.Equal(t, 3, len(books))
	}

	s = NewBookService(newMockBookDAOEmpty())
	books, err = s.FindAll()
	assert.NotNil(t, err)
	assert.Empty(t, books)
}

func newMockBookDAO() bookDAO {
	return &mockBookDAO{
		records: []models.Book{
			{Model: models.Model{ID: 1}, Title: "The Go Programming Language", CoverPictureURL: "https://www.gopl.io/cover.png", URL: "https://www.gopl.io/", AlternativeText: "The Go Programming Language"},
			{Model: models.Model{ID: 2}, Title: "Clean Code", CoverPictureURL: "https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1436202607i/3735293._SX318_.jpg", URL: "https://www.goodreads.com/book/show/3735293-clean-code", AlternativeText: "Clean Code: A Handbook of Agile Software Craftsmanship"},
			{Model: models.Model{ID: 3}, Title: "Software Craftsmanship", CoverPictureURL: "https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1370897661i/18054154.jpg", URL: "https://www.goodreads.com/book/show/18054154-software-craftsmanship", AlternativeText: "The Software Craftsman: Professionalism, Pragmatism, Pride"},
		},
	}
}

func newMockBookDAOEmpty() bookDAO {
	return &mockBookDAO{
		records: []models.Book{},
	}
}

func (m *mockBookDAO) FindAll() []models.Book {
	books := m.records
	return books
}

type mockBookDAO struct {
	records []models.Book
}
