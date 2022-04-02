package daos

import (
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/config"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPostDAO_Get(t *testing.T) {
	config.Config.DB = test_data.ResetDB()
	dao := NewPostDAO()

	post, err := dao.Get(1)

	expected := map[string]string{"Title": "First Blog Post", "Text": "This is blog about something."}

	assert.Nil(t, err)
	assert.Equal(t, expected["Title"], post.Title)
	assert.Equal(t, expected["Text"], post.Text)
	assert.Equal(t, 3, len(post.Tags))
	assert.Equal(t, 2, len(post.Sections))
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

func TestPostDAO_FindAllOrdered(t *testing.T) {
	timeFormat := "2006-01-02 15:04:05"
	config.Config.DB = test_data.ResetDB()
	dao := NewPostDAO()

	posts := dao.FindAll()

	firstPostDate, _ := time.Parse(timeFormat, "2019-05-30 19:00:00")
	assert.Equal(t, firstPostDate, posts[0].PostedOn)
	secondPostDate, _ := time.Parse(timeFormat, "2019-02-24 13:00:00")
	assert.Equal(t, secondPostDate, posts[1].PostedOn)
	thirdPostDate, _ := time.Parse(timeFormat, "2018-08-24 14:00:00")
	assert.Equal(t, thirdPostDate, posts[2].PostedOn)
}

func TestPostDAO_FindAllByTag(t *testing.T) {
	config.Config.DB = test_data.ResetDB()
	dao := NewPostDAO()

	posts := dao.FindAllByTag("python")

	assert.Equal(t, 2, len(posts))

	expected := map[string]string{"Title": "3rd Blog Post", "Text": "Another dummy content"}
	post := dao.FindAllByTag("vue")

	assert.Equal(t, 1, len(post))
	assert.Equal(t, expected["Title"], post[0].Title)
	assert.Equal(t, expected["Text"], post[0].Text)
}
