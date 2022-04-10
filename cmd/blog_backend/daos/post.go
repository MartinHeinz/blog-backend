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
	config.Config.DB.Order("posted_on DESC").
		Find(&posts)
	return posts
}

func (dao *PostDAO) FindAllByTag(tagName string) []models.Post {
	var posts []models.Post
	config.Config.DB.
		Select("title", "text", "posts.id", "posted_on", "next_post_id", "previous_post_id", "author").
		Joins("JOIN tags ON tags.post_id = posts.id").
		Where("lower(name) = lower(?) and post_id is not null", tagName).
		Order("posted_on DESC").
		Find(&posts)

	return posts
}
