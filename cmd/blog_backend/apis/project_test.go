package apis

import (
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/test_data"
	"net/http"
	"testing"
)

func TestProject(t *testing.T) {
	path := test_data.GetTestCaseFolder()
	runAPITests(t, []apiTestCase{
		{"t1 - get all projects", "GET", "/projects/", "/projects/", "", GetProjects, http.StatusOK, path + "/project_t1.json"},
	})
}
