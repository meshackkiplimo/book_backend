package main

import (
	"log"

	"bookstore/internal/database"
	"bookstore/internal/models"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading from system environment variables.")
	}

	database.Connect()

	books := []models.Book{
		{Title: "The Go Programming Language", Author: "Alan A. A. Donovan", Price: 39.99, Description: "A comprehensive guide to Go."},
		{Title: "Clean Code", Author: "Robert C. Martin", Price: 45.50, Description: "A book on software craftsmanship."},
		{Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", Price: 29.99, Description: "An epic fantasy novel."},
	}

	for _, book := range books {
		result := database.DB.Create(&book)
		if result.Error != nil {
			log.Printf("Failed to seed book '%s': %v", book.Title, result.Error)
		} else {
			log.Printf("Seeded book: %s", book.Title)
		}
	}
	log.Println("Database seeding complete.")
}
