package services

import (
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSectionService(t *testing.T) {
	dao := newMockSectionDAO()
	s := NewSectionService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestSectionService_FindAll(t *testing.T) {
	s := NewSectionService(newMockSectionDAO())
	sections, err := s.FindAll(1)
	if assert.Nil(t, err) && assert.NotEmpty(t, sections) {
		assert.Equal(t, 3, len(sections))
	}

	sections, err = s.FindAll(100)
	assert.NotNil(t, err)
	assert.Empty(t, sections)
}

func newMockSectionDAO() sectionDAO {
	return &mockSectionDAO{
		records: []models.Section{
			{Model: models.Model{ID: 1}, PostID: 1, Name: "Title"},
			{Model: models.Model{ID: 2}, PostID: 2, Name: "Title"},
			{Model: models.Model{ID: 3}, PostID: 3, Name: "Title"},
			{Model: models.Model{ID: 4}, PostID: 1, Name: "Subsection"},
			{Model: models.Model{ID: 5}, PostID: 1, Name: "Subsection 2"},
		},
	}
}

func (m *mockSectionDAO) FindAll(postId uint) []models.Section {
	var sections []models.Section

	for _, record := range m.records {
		if record.PostID == postId {
			sections = append(sections, record)
		}
	}
	return sections
}

type mockSectionDAO struct {
	records []models.Section
}
