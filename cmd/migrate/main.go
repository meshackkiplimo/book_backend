package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading from system environment variables.")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL not set in .env file or environment.")
	}

	m, err := migrate.New(
		"file://migrations",
		databaseURL)
	if err != nil {
		log.Fatalf("Could not create migration instance: %v", err)
	}

	cmd := os.Args[1]
	switch cmd {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Migration failed: %v", err)
		}
		log.Println("Migrations applied successfully.")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Migration rollback failed: %v", err)
		}
		log.Println("Migrations rolled back successfully.")
	default:
		log.Fatalf("Unknown command: %s. Use 'up' or 'down'.", cmd)
	}
}
