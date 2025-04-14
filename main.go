package main

import (
	"fmt"
	"log"
	"os"

	"go-bookcase/config"
	"go-bookcase/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Memuat file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := config.SetupDatabase()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	r := gin.Default()

	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
