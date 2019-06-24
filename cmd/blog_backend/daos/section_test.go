package daos

import (
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/config"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSectionDAO_FindAll(t *testing.T) {
	config.Config.DB = test_data.ResetDB()
	dao := NewSectionDAO()

	sections := dao.FindAll(2)

	assert.Equal(t, 4, len(sections))
}

func TestSectionDAO_FindEmpty(t *testing.T) {
	config.Config.DB = test_data.ResetDB()
	dao := NewSectionDAO()

	sections := dao.FindAll(9999)

	assert.Empty(t, sections)
}
