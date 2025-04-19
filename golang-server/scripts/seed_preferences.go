package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	fmt.Println("Connected to database successfully")

	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	var count int
	err = tx.QueryRow("SELECT COUNT(*) FROM preference_metadata WHERE type = 'vibe' AND key IN ('Work', 'Study', 'Chill')").Scan(&count)
	if err != nil {
		log.Fatalf("Failed to query existing vibe preferences: %v", err)
	}

	if count == 0 {
		// Insert vibe preferences
		vibePreferences := []string{"Work", "Study", "Chill"}
		for _, vibe := range vibePreferences {
			_, err = tx.Exec(
				"INSERT INTO preference_metadata (id, key, type, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW())",
				uuid.New(), vibe, "vibe",
			)
			if err != nil {
				log.Fatalf("Failed to insert vibe preference '%s': %v", vibe, err)
			}
			fmt.Printf("Inserted vibe preference: %s\n", vibe)
		}
	} else {
		fmt.Println("Vibe preferences already exist, skipping insertion")
	}

	// Check if style entries already exist
	err = tx.QueryRow("SELECT COUNT(*) FROM preference_metadata WHERE type = 'style' AND key IN ('Modern', 'Vintage', 'Minimalist')").Scan(&count)
	if err != nil {
		log.Fatalf("Failed to query existing style preferences: %v", err)
	}

	if count == 0 {
		// Insert style preferences
		stylePreferences := []string{"Modern", "Vintage", "Minimalist"}
		for _, style := range stylePreferences {
			_, err = tx.Exec(
				"INSERT INTO preference_metadata (id, key, type, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW())",
				uuid.New(), style, "style",
			)
			if err != nil {
				log.Fatalf("Failed to insert style preference '%s': %v", style, err)
			}
			fmt.Printf("Inserted style preference: %s\n", style)
		}
	} else {
		fmt.Println("Style preferences already exist, skipping insertion")
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}

	fmt.Println("Successfully seeded preferences_metadata table")
}
