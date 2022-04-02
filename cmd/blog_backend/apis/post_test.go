package apis

import (
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/test_data"
	"net/http"
	"testing"
)

func TestPost(t *testing.T) {
	path := test_data.GetTestCaseFolder()
	runAPITests(t, []apiTestCase{
		{"t1 - get a Post", "GET", "/posts/by-id/:id", "/posts/by-id/1", "", GetPost, http.StatusOK, path + "/post_t1.json"},
		{"t2 - get a Post not Present", "GET", "/posts/by-id/:id", "/posts/by-id/9999", "", GetPost, http.StatusNotFound, ""},
		{"t3 - get all posts", "GET", "/posts/", "/posts/", "", GetPosts, http.StatusOK, path + "/post_t3.json"},
		{"t4 - get all posts by tag name", "GET", "/posts/by-tag/:tag_name", "/posts/by-tag/python", "", GetPostsByTag, http.StatusOK, path + "/post_t4.json"},
	})
}
