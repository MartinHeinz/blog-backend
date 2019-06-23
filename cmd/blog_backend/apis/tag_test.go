package apis

import (
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/test_data"
	"net/http"
	"testing"
)

func TestTag(t *testing.T) {
	path := test_data.GetTestCaseFolder()
	runAPITests(t, []apiTestCase{
		{"t1 - get all tags for post", "GET", "/tags/:post_id", "/tags/1", "", GetTags, http.StatusOK, path + "/tag_t1.json"},
		{"t2 - try to get tags, that are not present", "GET", "/tags/:post_id", "/tags/9999", "", GetTags, http.StatusNotFound, ""},
	})
}
