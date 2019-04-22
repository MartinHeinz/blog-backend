package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetData(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/testGetData", testGetData)

	body := gin.H{
		"message": "Get Hello",
	}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/testGetData", nil)
	router.ServeHTTP(w, r)

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, body["message"], response["message"])
}
