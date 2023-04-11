package main

import (
	"database/sql"
	"fmt"
	router2 "github.com/phuc-create/go-simple-crud/router"
	"log"
	"net/http"
	"os"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUsername, dbPassword, dbName)
	db, err := sql.Open("postgres", connStr)

	//if err != nil {
	//	log.Fatal(err)
	//}
	router, err := router2.NewRouter(db)
	if err != nil {
		log.Fatalf("error initializing router: %v", err)
	}

	log.Fatal(http.ListenAndServe(":3000", router))
}
