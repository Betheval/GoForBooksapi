package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title,omitempty"`
	Author    string             `bson:"author,omitempty"`
	ISBN      string             `bson:"isbn,omitempty"`
	Publisher string             `bson:"publisher,omitempty"`
	Year      int                `bson:"year,omitempty"`
}

var client *mongo.Client

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/books", getBooks)

	r.GET("/book/id/:id", func(c *gin.Context) {
		//TODO: return a book by ID
	})

	r.GET("/book/isbn/:isbn", func(c *gin.Context) {
		//TODO: return a book by ISBN
	})

	r.DELETE("/books/:id", func(c *gin.Context) {
		//TODO: delete a book by ID
	})

	r.POST("/books", func(c *gin.Context) {
		//TODO: Add a book
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, " ")
}

//MongoDB connection

func connectToMongoDB() *mongo.Client {
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("MONGODB_URI not set in environment")
	}
	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}
