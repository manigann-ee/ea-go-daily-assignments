package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/hello", nil)

	helloHandler(recorder, request)

	responseBody := recorder.Result().Body
	response, _ := io.ReadAll(responseBody)

	assert.Equal(t, "hello", string(response))
}
