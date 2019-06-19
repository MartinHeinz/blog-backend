package services

import (
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/models"
)

type postDAO interface {
	Get(id string) (*models.Post, error)
}

type PostService struct {
	dao postDAO
}

// NewPostService creates a new PostService with the given post DAO.
func NewPostService(dao postDAO) *PostService {
	return &PostService{dao}
}

func (s *PostService) Get(id string) (*models.Post, error) {
	return s.dao.Get(id)
}
