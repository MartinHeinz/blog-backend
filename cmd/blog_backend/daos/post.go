package daos

import (
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/config"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/models"
)

// PostDAO persists post data in database
type PostDAO struct{}

// NewPostDAO creates a new PostDAO
func NewPostDAO() *PostDAO {
	return &PostDAO{}
}

func (dao *PostDAO) Get(id uint) (*models.Post, error) {
	var post models.Post
	err := config.Config.DB.Where("id = ?", id).
		Preload("Tags").
		Preload("Sections").
		First(&post).
		Error

	return &post, err
}

func (dao *PostDAO) FindAll() []models.Post {
	var posts []models.Post
	config.Config.DB.Find(&posts)
	return posts
}
