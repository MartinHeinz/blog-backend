package apis

import (
	"net/http"
	"testing"
)

func TestPost(t *testing.T) {
	runAPITests(t, []apiTestCase{
		{"t1 - get a Post", "GET", "/posts/:id", "/posts/1", "", GetPost, http.StatusOK, map[string]string{"Title": "First Blog Post", "Text": "This is blog about something."}},
		{"t2 - get a Post not Present", "GET", "/posts/:id", "/posts/9999", "", GetPost, http.StatusNotFound, nil},
	})
}
