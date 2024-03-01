package server

import (
	"api-gin/src/config"
	"api-gin/src/db/mongo"
	"api-gin/src/routes"
	"context"
	"fmt"
	"log"
)

func Init() {

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}
	port := cfg.AppPort
	if port == "" {
		port = "8080"
	}

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
