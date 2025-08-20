package api

import (
	"bookstore/internal/database"
	"bookstore/internal/models"
	"encoding/json"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r http.Request) {
	var books []models.Book
	database.DB.Find(&books)
	json.NewEncoder(w).Encode(books)
}
