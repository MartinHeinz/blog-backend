package daos

import (
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/config"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostDAO_Get(t *testing.T) {
	config.Config.DB = test_data.ResetDB()
	dao := NewPostDAO()

	post, err := dao.Get(1)

	expected := map[string]string{"Title": "First Blog Post", "Text": "This is blog about something."}

	assert.Nil(t, err)
	assert.Equal(t, expected["Title"], post.Title)
	assert.Equal(t, expected["Text"], post.Text)
}

func TestPostDAO_GetNotPresent(t *testing.T) {
	config.Config.DB = test_data.ResetDB()
	dao := NewPostDAO()

	post, err := dao.Get(9999)

	assert.NotNil(t, err)
	assert.Equal(t, "", post.Title)
	assert.Equal(t, "", post.Text)
}

func TestPostDAO_FindAll(t *testing.T) {
	config.Config.DB = test_data.ResetDB()
	dao := NewPostDAO()

	posts := dao.FindAll()

	assert.Equal(t, 3, len(posts))
}
