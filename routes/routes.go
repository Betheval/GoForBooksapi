package routes

import (
	"booksapi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/books", controllers.GetBooks)
	r.GET("/book/id/:id", controllers.GetBookByID)
	r.GET("/book/isbn/:isbn", controllers.GetBookByISBN)
	r.POST("/book", controllers.AddBook)
	r.DELETE("/book/deletion/:id", controllers.DeleteBookByID)

	return r
}
