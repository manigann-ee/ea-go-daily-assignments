package songs

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

var r *gin.Engine

func init() {
	r = gin.New()
	r.POST("/api/songs", CreateSongHandler)
	r.GET("/api/songs/:id", GetSongHandler)
	r.GET("/api/songs", GetSongsHandler)
}

func TestCreateSongHandler(t *testing.T) {
	reqBody := `{
		"title": "title one",
        "singer": "singer one",
		"writer": "writer one",
		"releasedDate": "2022-12-22T00:00:00Z"
	}`

	req, _ := http.NewRequest("POST", "/api/songs", strings.NewReader(reqBody))
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	responseBody, _ := io.ReadAll(recorder.Body)
	response := SongResponse{}
	err := json.Unmarshal(responseBody, &response)
	assert.Nil(t, err, "expected error to be nil")
	assert.Equal(t, 1, response.Id)
	assert.Equal(t, "title one", response.Title)
	assert.Equal(t, "singer one", response.Singer)
}

func TestCreateSongHandler_ReturnsErrorWhenSongRequestIsInvalid(t *testing.T) {
	reqBody := `{
		"title": "title one",
        "singer": "singer one",
		"writer": "writer one",
		"releasedDate": "2022-12-22"
	}`

	req, _ := http.NewRequest("POST", "/api/songs", strings.NewReader(reqBody))
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestGetSongHandlerSuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/songs/1", nil)
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	responseBody, _ := io.ReadAll(recorder.Body)
	response := SongResponse{}
	err := json.Unmarshal(responseBody, &response)

	assert.Nil(t, err, "expected error to be nil")
	assert.Equal(t, 1, response.Id)
	assert.Equal(t, "Title One", response.Title)
	assert.Equal(t, "Singer One", response.Singer)
}

func TestGetSongsHandlerSuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/songs", nil)
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	responseBody, _ := io.ReadAll(recorder.Body)
	response := QueryResponse{}
	err := json.Unmarshal(responseBody, &response)

	assert.Nil(t, err, "expected error to be nil")
	assert.Equal(t, 2, response.Total)
	assert.Equal(t, "title one", response.Songs[0].Title)
	assert.Equal(t, "singer one", response.Songs[0].Singer)
	assert.Equal(t, "title two", response.Songs[1].Title)
	assert.Equal(t, "singer two", response.Songs[1].Singer)
}
