package middleware

import (
	"fmt"
	brotli "github.com/anargu/gin-brotli"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestBrotliMiddleware(t *testing.T) {
	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(resp)

	r.Use(brotli.Brotli(brotli.DefaultCompression))

	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("World at %s", time.Now()))
	})
	c.Request, _ = http.NewRequest(http.MethodGet, "/test", nil)
	c.Request.Header.Set("Content-Type", "application/json")
	c.Request.Header.Set("Accept-Encoding", "br")
	r.ServeHTTP(resp, c.Request)

	log.Println(resp.Header())
	assert.Equal(t, "br", resp.Header().Get("Content-Encoding"))
}
