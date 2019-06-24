package daos

import (
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/config"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/models"
)

// TagDAO persists tag data in database
type SectionDAO struct{}

// NewSectionDAO creates a new SectionDAO
func NewSectionDAO() *SectionDAO {
	return &SectionDAO{}
}

func (dao *SectionDAO) FindAll(postId uint) []models.Section {
	var sections []models.Section
	config.Config.DB.Where("post_id = ?", postId).Find(&sections)
	return sections
}
