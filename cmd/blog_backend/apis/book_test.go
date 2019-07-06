package apis

import (
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/test_data"
	"net/http"
	"testing"
)

func TestBook(t *testing.T) {
	path := test_data.GetTestCaseFolder()
	runAPITests(t, []apiTestCase{
		{"t1 - get all books", "GET", "/books/", "/books/", "", GetBooks, http.StatusOK, path + "/book_t1.json"},
	})
}
