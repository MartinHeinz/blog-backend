package daos

import (
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/config"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/models"
)

// TagDAO persists tag data in database
type TagDAO struct{}

// NewTagDAO creates a new TagDAO
func NewTagDAO() *TagDAO {
	return &TagDAO{}
}

func (dao *TagDAO) FindAll(postId uint) []models.Tag {
	var tags []models.Tag
	config.Config.DB.Where("post_id = ?", postId).Find(&tags)
	return tags
}
