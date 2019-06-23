package daos

import (
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/config"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTagDAO_Get(t *testing.T) {
	config.Config.DB = test_data.ResetDB()
	dao := NewTagDAO()

	tags := dao.FindAll(1)

	assert.Equal(t, 3, len(tags))
}

func TestTagDAO_FindEmpty(t *testing.T) {
	config.Config.DB = test_data.ResetDB()
	dao := NewTagDAO()

	tags := dao.FindAll(9999)

	assert.Empty(t, tags)
}
