package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {

	// Load .env only for local development
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Render PostgreSQL requires SSL
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		host,
		port,
		user,
		password,
		dbname,
	)

	var err error

	// Initialize the global DB variable
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("✅ Connected to PostgreSQL")
}
