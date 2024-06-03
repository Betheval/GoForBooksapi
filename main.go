package main

import (
	"booksapi/config"
	"booksapi/routes"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func main() {
	config.ConnectToMongoDB()
	defer config.Client.Disconnect(context.Background())

	r := routes.SetupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
