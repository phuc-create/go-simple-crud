package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/phuc-create/go-simple-crud/internal/repository"
	"github.com/phuc-create/go-simple-crud/internal/repository/controllers/users"
	router2 "github.com/phuc-create/go-simple-crud/router"
	"log"
	"os"
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
		"host=%s port=%s users=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUsername,
		dbPassword,
		dbName,
	)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	router2.New(ctx, db)
}

type controllers struct {
	usersController users.Controllers
}

func initControllers(db *sql.DB) controllers {
	repo := repository.New(db)
	return controllers{usersController: users.New(repo)}
}
