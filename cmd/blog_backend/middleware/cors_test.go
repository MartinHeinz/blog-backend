package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCORSMiddlewareGET(t *testing.T) {
	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(resp)

	r.Use(CORSMiddleware())

	r.GET("/test", func(c *gin.Context) {
		c.Status(200)
	})
	c.Request, _ = http.NewRequest(http.MethodGet, "/test", nil)
	r.ServeHTTP(resp, c.Request)

	assert.Equal(t, "*", resp.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "true", resp.Header().Get("Access-Control-Allow-Credentials"))
	assert.Equal(t, "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With",
		resp.Header().Get("Access-Control-Allow-Headers"))
	assert.Equal(t, "POST, OPTIONS, GET, PUT", resp.Header().Get("Access-Control-Allow-Methods"))
}

func TestCORSMiddlewareOPTIONS(t *testing.T) {
	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(resp)

	r.Use(CORSMiddleware())

	r.OPTIONS("/test", func(c *gin.Context) {})

	c.Request, _ = http.NewRequest(http.MethodOptions, "/test", nil)
	r.ServeHTTP(resp, c.Request)

	assert.Equal(t, http.StatusNoContent, resp.Code)
}
