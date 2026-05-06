package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}


	connStr := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal("Unable to connect to DB:", err)
	}
	defer conn.Close(context.Background())

	log.Println("Connected to DB")

	// bootstrap
	
	runScraper(conn)
	// runAPI(conn)
}

func runScraper(conn *pgx.Conn) {
	log.Println("Running scraper...")

	// Pagaidām fake data (testam)
	_, err := conn.Exec(context.Background(),
		`INSERT INTO internships (title, company, location, url)
		 VALUES ($1, $2, $3, $4)
		 ON CONFLICT (url) DO NOTHING`,
		"Backend Intern",
		"Test Company",
		"Remote",
		"https://example.com/job/1",
	)

	if err != nil {
		log.Println("Insert error:", err)
	}

	log.Println("Scraper finished")
}