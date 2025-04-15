package main

import (
	"fmt"
	"log"
	"os"

	"go-bookcase/config"
	"go-bookcase/routes"

	_ "go-bookcase/docs" // Mengimpor package docs untuk Swagger

	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	// Memuat file .env
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	db, err := config.SetupDatabase()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	r := gin.Default()

	// Setup Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Setup routes
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Jalankan aplikasi
	err = r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
