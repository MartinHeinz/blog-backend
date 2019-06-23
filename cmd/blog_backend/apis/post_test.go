package apis

import (
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/test_data"
	"net/http"
	"testing"
)

func TestPost(t *testing.T) {
	path := test_data.GetTestCaseFolder()
	runAPITests(t, []apiTestCase{
		{"t1 - get a Post", "GET", "/posts/:id", "/posts/1", "", GetPost, http.StatusOK, path + "/post_t1.json"},
		{"t2 - get a Post not Present", "GET", "/posts/:id", "/posts/9999", "", GetPost, http.StatusNotFound, ""},
	})
}
