package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_createBooksHandler(t *testing.T) {
	engine := gin.New()
	engine.POST("/books", createBooksHandler)

	requestBody := `{
		"Id": 1,
		"Title": "BookOne",
		"Price": 50
	}`
	request, _ := http.NewRequest("POST", "/books", strings.NewReader(requestBody))
	recorder := httptest.NewRecorder()

	engine.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusCreated, recorder.Code)
}

func Test_getBooksHandler(t *testing.T) {
	engine := gin.New()
	engine.POST("/books", createBooksHandler)
	engine.GET("/books", getBooksHandler)

	requestBody := `{
		"Id": 1,
		"Title": "BookOne",
		"Price": 50
	}`
	http.NewRequest("POST", "/books", strings.NewReader(requestBody))

	request, _ := http.NewRequest("GET", "/books", nil)
	recorder := httptest.NewRecorder()

	engine.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	responseBody, _ := io.ReadAll(recorder.Body)
	var response []Book
	err := json.Unmarshal(responseBody, &response)
	assert.Nil(t, err, "expected unmarshal error to be nil")
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, 1, response[0].Id)
	assert.Equal(t, "BookOne", response[0].Title)
	assert.Equal(t, float64(50), response[0].Price)
}
