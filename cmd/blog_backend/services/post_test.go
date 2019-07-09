package services

import (
	"errors"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPostService(t *testing.T) {
	dao := newMockPostDAO()
	s := NewPostService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestPostService_Get(t *testing.T) {
	s := NewPostService(newMockPostDAO())
	post, err := s.Get(1)
	if assert.Nil(t, err) && assert.NotNil(t, post) {
		assert.Equal(t, "Test Title", post.Title)
	}

	post, err = s.Get(100)
	assert.NotNil(t, err)
}

func TestPostService_FindAll(t *testing.T) {
	s := NewPostService(newMockPostDAO())
	books, err := s.FindAll()
	if assert.Nil(t, err) && assert.NotEmpty(t, books) {
		assert.Equal(t, 2, len(books))
	}

	s = NewPostService(newMockPostDAOEmpty())
	books, err = s.FindAll()
	assert.NotNil(t, err)
	assert.Empty(t, books)
}

func newMockPostDAO() postDAO {
	return &mockPostDAO{
		records: []models.Post{
			{Model: models.Model{ID: 1}, Title: "Test Title", Text: "Test Text."},
			{Model: models.Model{ID: 2}, Title: "Test Title 2", Text: "Test Text 2."},
		},
	}
}

func newMockPostDAOEmpty() postDAO {
	return &mockPostDAO{
		records: []models.Post{},
	}
}

func (m *mockPostDAO) Get(id uint) (*models.Post, error) {
	for _, record := range m.records {
		if record.ID == id {
			return &record, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *mockPostDAO) FindAll() []models.Post {
	posts := m.records
	return posts
}

type mockPostDAO struct {
	records []models.Post
}
