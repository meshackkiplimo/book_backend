package main

import (
	"bookstore/internal/api"
	"bookstore/internal/database"
	"bookstore/internal/models"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env file found, skipping loading")
	}
	// Connect to the database
	database.Connect()

	database.DB.AutoMigrate(&models.Book{})

	router := api.Routes()
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000" // Default port if not set
	}
	log.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
