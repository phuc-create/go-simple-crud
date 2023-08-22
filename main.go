package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	router2 "github.com/phuc-create/go-simple-crud/router"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

func main() {
	ctx := context.Background()
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
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

	if err := db.Ping(); err != nil {
		db.Close()
	}

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB already connected")
	fmt.Println("this is commit 1")
	fmt.Println("this is commit 2")
	fmt.Println("this is commit 3")
	fmt.Println("this is commit 456")
	fmt.Println("this is commit 789")
	router2.New(ctx, db)
}
