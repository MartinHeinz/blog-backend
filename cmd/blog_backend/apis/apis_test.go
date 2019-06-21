package apis

import (
	"bytes"
	"encoding/json"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/config"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/test_data"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type apiTestCase struct {
	tag        string
	method     string
	urlToServe string
	urlToHit   string
	body       string
	function   gin.HandlerFunc
	status     int
	response   map[string]string
}

func newRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	config.Config.DB = test_data.ResetDB()

	return router
}

func testAPI(router *gin.Engine, method string, urlToServe string, urlToHit string, function gin.HandlerFunc, body string) *httptest.ResponseRecorder {
	router.Handle(method, urlToServe, function)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(method, urlToHit, bytes.NewBufferString(body))
	router.ServeHTTP(res, req)
	return res
}

func runAPITests(t *testing.T, tests []apiTestCase) {
	for _, test := range tests {
		router := newRouter()
		res := testAPI(router, test.method, test.urlToServe, test.urlToHit, test.function, test.body)
		assert.Equal(t, test.status, res.Code, test.tag)
		if test.response != nil {
			var jsonRes map[string]string
			json.Unmarshal([]byte(res.Body.String()), &jsonRes)
			for key, val := range test.response {
				assert.Equal(t, val, jsonRes[key], test.tag)
			}
		}
	}
}
