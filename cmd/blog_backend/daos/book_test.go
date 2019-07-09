package daos

import (
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/config"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookDAO_FindAll(t *testing.T) {
	config.Config.DB = test_data.ResetDB()
	dao := NewBookDAO()

	books := dao.FindAll()

	assert.Equal(t, 4, len(books))
}
