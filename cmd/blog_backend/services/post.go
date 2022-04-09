package services

import (
	"context"
	"fmt"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/models"
)

type postDAO interface {
	Get(ctx context.Context, id uint) (*models.Post, error)
	FindAll() []models.Post
	FindAllByTag(tagName string) []models.Post
}

type PostService struct {
	dao postDAO
}

// NewPostService creates a new PostService with the given post DAO.
func NewPostService(dao postDAO) *PostService {
	return &PostService{dao}
}

func (s *PostService) Get(ctx context.Context, id uint) (*models.Post, error) {
	return s.dao.Get(ctx, id)
}

func (s *PostService) FindAll() ([]models.Post, error) {
	posts := s.dao.FindAll()
	err := fmt.Errorf("no posts found")
	if len(posts) > 0 {
		return posts, nil
	}
	return posts, err
}

func (s *PostService) FindAllByTag(tagName string) ([]models.Post, error) {
	posts := s.dao.FindAllByTag(tagName)
	err := fmt.Errorf("no posts found")
	if len(posts) > 0 {
		return posts, nil
	}
	return posts, err
}
