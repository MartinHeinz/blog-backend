package apis

import (
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/test_data"
	"net/http"
	"testing"
)

func TestSection(t *testing.T) {
	path := test_data.GetTestCaseFolder()
	runAPITests(t, []apiTestCase{
		{"t1 - get all sections for post", "GET", "/sections/:post_id", "/sections/1", "", GetSections, http.StatusOK, path + "/section_t1.json"},
		{"t2 - try to get sections, that are not present", "GET", "/sections/:post_id", "/sections/9999", "", GetSections, http.StatusNotFound, ""},
	})
}
