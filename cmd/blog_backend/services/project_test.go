package services

import (
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProjectService(t *testing.T) {
	dao := newMockProjectDAO()
	s := NewProjectService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestProjectService_FindAll(t *testing.T) {
	s := NewProjectService(newMockProjectDAO())
	projects, err := s.FindAll()
	if assert.Nil(t, err) && assert.NotEmpty(t, projects) {
		assert.Equal(t, 2, len(projects))
	}

	s = NewProjectService(newMockProjectDAOEmpty())
	projects, err = s.FindAll()
	assert.NotNil(t, err)
	assert.Empty(t, projects)
}

func newMockProjectDAO() projectDAO {
	return &mockProjectDAO{
		records: []models.Project{
			{Model: models.Model{ID: 1}, Name: "IoT Cloud", ThumbnailPictureURL: "https://via.placeholder.com/150", URL: "https://github.com/MartinHeinz/IoT-Cloud", Description: "Cloud framework for IoT (Internet of Things), which focuses on security and privacy of its users, their devices and data"},
			{Model: models.Model{ID: 2}, Name: "Blog & Personal Website", ThumbnailPictureURL: "https://via.placeholder.com/150", URL: "https://github.com/MartinHeinz/blog-backend", Description: "This website. Goal of this project was to learn Go and Vue.js and as a byproduct I created personal website and blog."},
		},
	}
}

func newMockProjectDAOEmpty() projectDAO {
	return &mockProjectDAO{
		records: []models.Project{},
	}
}

func (m *mockProjectDAO) FindAll() []models.Project {
	projects := m.records
	return projects
}

type mockProjectDAO struct {
	records []models.Project
}
