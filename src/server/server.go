package server

import (
	"api-gin/src/routes"
	"api-gin/src/db"
	"context"
	"log"
)

func Init() {
	port := "8080"

	// Connect to MongoDB
	client, err := db.ConnectToMongoDB()
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	// Start the server on the configured port
	r := routes.NewRouter()
	r.Run(":" + port)

}
