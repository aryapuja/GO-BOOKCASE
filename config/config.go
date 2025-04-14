package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// SetupDatabase initializes the database connection
func SetupDatabase() (*sql.DB, error) {
	// Ambil variabel lingkungan dari file .env
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Cek apakah variabel lingkungan sudah terisi
	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		return nil, fmt.Errorf("DATA KONEKSI DATABASE ADA YANG BELUM DI SET")
	}

	// Format URL koneksi PostgreSQL
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Buka koneksi ke database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Tes koneksi ke database
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to the database")
	return db, nil
}
