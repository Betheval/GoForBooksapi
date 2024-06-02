package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/books", func(c *gin.Context) {
		//TODO: return first 100 books
	})

	r.GET("/books:id", func(c *gin.Context) {
		//TODO: return a book by ID
	})

	r.GET("/books:isbn", func(c *gin.Context) {
		//TODO: return a book by ISBN
	})

	r.DELETE("/books:id", func(c *gin.Context) {
		//TODO: delete a book by ID
	})

	r.POST("/books", func(c *gin.Context) {
		//TODO: Add a book
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

//modify a book
