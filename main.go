package main

import (
	"database/sql"
	"fmt"
	godotenv "github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/phuc-create/go-simple-crud/controllers/user"
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
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUsername, dbPassword, dbName)
	fmt.Println("Connection string: ", connStr)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	userController := user.New(db)
	router2.New(db, userController)
}
