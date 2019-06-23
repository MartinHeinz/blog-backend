package services

import (
	"fmt"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/models"
)

type tagDAO interface {
	FindAll(id uint) []models.Tag
}

type TagService struct {
	dao tagDAO
}

// NewTagService creates a new TagService with the given tag DAO.
func NewTagService(dao tagDAO) *TagService {
	return &TagService{dao}
}

func (s *TagService) FindAll(postId uint) ([]models.Tag, error) {
	tags := s.dao.FindAll(postId)
	err := fmt.Errorf("no tags found for post_id: %v", postId)
	if len(tags) > 0 {
		return tags, nil
	}
	return tags, err
}
