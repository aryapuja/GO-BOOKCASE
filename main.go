package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-bookcase/config"
	"go-bookcase/routes"
)

func main() {
	// Memuat file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Setup koneksi database
	db, err := config.SetupDatabase()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	// Initialize Gin router
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r)

	// Start the server on the specified port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
