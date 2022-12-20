package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Book struct {
	Id    int
	Title string
	Price float64
}

var books []Book

func main() {
	r := gin.Default()
	r.POST("/books", createBooksHandler)
	r.GET("/books", getBooksHandler)
	r.Run()
}

func createBooksHandler(context *gin.Context) {
	book := Book{}
	context.BindJSON(&book)
	books = append(books, book)
	context.JSON(http.StatusCreated, &books)
}

func getBooksHandler(c *gin.Context) {
	c.JSON(http.StatusOK, &books)
}
