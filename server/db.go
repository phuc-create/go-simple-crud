package server

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func DBConnection() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	fmt.Println("DB already connected")
	return db, nil
}
