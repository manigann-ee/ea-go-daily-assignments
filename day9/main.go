package main

import (
	"github.com/gin-gonic/gin"
	"songs/pkg/songs"
)

func main() {
	engine := gin.New()

	engine.Use(gin.Recovery())

	engine.POST("/api/songs", songs.CreateSongHandler)
	engine.GET("/api/songs/:id", songs.GetSongHandler)
	engine.GET("/api/songs", songs.GetSongsHandler)

	engine.Run(":5050")
}
