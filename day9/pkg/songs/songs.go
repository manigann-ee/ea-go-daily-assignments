package songs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type SongRequest struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	Singer       string    `json:"singer"`
	Writer       string    `json:"writer"`
	ReleasedDate time.Time `json:"releasedDate"`
}

type SongResponse struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	Singer       string    `json:"singer"`
	Writer       string    `json:"writer"`
	ReleasedDate time.Time `json:"releasedDate"`
}

type QueryResponse struct {
	Songs []SongResponse `json:"songs"`
	Total int            `json:"total"`
}

func CreateSongHandler(ctx *gin.Context) {
	songRequest := SongRequest{}
	err := ctx.BindJSON(&songRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to parse request",
		})
	}

	songResponse := SongResponse{
		Id:           1,
		Title:        songRequest.Title,
		Singer:       songRequest.Singer,
		Writer:       songRequest.Writer,
		ReleasedDate: songRequest.ReleasedDate,
	}

	ctx.JSON(http.StatusCreated, songResponse)
}

func GetSongHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	songId, _ := strconv.Atoi(id)

	response := SongResponse{
		Id:           songId,
		Title:        "Title One",
		Singer:       "Singer One",
		Writer:       "Writer One",
		ReleasedDate: time.Time{},
	}

	ctx.JSON(http.StatusOK, response)
}

func GetSongsHandler(ctx *gin.Context) {
	title := ctx.Query("title")
	fmt.Printf("title is: %s", title)

	songResponses := []SongResponse{
		{
			Id:           1,
			Title:        "title one",
			Singer:       "singer one",
			Writer:       "writer one",
			ReleasedDate: time.Time{},
		},
		{
			Id:           2,
			Title:        "title two",
			Singer:       "singer two",
			Writer:       "writer two",
			ReleasedDate: time.Time{},
		},
	}

	queryResponse := QueryResponse{
		Songs: songResponses,
		Total: 2,
	}

	ctx.JSON(http.StatusOK, queryResponse)
}
