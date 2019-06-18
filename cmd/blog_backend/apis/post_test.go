package apis

import (
	"encoding/json"
	"fmt"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/config"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/models"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/test_data"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	config.Config.DB = test_data.ResetDB()

	var post models.Post
	config.Config.DB.First(&post)
	//os.Stderr.WriteString(fmt.Sprintf("%v", post.ID))

	path := "/posts/:id"
	router.GET(path, GetPost)

	body := gin.H{
		"title": "First Blog Post",
		"text":  "This is blog about something.",
	}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", fmt.Sprintf("/posts/%d", 1), nil)
	router.ServeHTTP(w, r)

	var response map[string]string
	json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, body["title"], response["Title"])
	assert.Equal(t, body["text"], response["Text"])
}

func TestGetPostNotPresent(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	path := "/posts/:id"
	router.GET(path, GetPost)

	w := httptest.NewRecorder()
	// Post with this ID not present
	r, _ := http.NewRequest("GET", fmt.Sprintf("/posts/%d", 9999), nil)
	router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
