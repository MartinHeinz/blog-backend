package services

import (
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTagService(t *testing.T) {
	dao := newMockTagDAO()
	s := NewTagService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestTagService_FindAll(t *testing.T) {
	s := NewTagService(newMockTagDAO())
	tags, err := s.FindAll(1)
	if assert.Nil(t, err) && assert.NotEmpty(t, tags) {
		assert.Equal(t, 2, len(tags))
	}

	tags, err = s.FindAll(100)
	assert.NotNil(t, err)
	assert.Empty(t, tags)
}

func newMockTagDAO() tagDAO {
	return &mockTagDAO{
		records: []models.Tag{
			{Model: models.Model{ID: 1}, PostID: 1, Name: "Python"},
			{Model: models.Model{ID: 2}, PostID: 1, Name: "Golang"},
			{Model: models.Model{ID: 3}, PostID: 3, Name: "Crypto"},
			{Model: models.Model{ID: 4}, PostID: 2, Name: "Python"},
		},
	}
}

func (m *mockTagDAO) FindAll(postId uint) []models.Tag {
	var tags []models.Tag

	for _, record := range m.records {
		if record.PostID == postId {
			tags = append(tags, record)
		}
	}
	return tags
}

type mockTagDAO struct {
	records []models.Tag
}
