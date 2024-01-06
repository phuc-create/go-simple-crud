package main

import (
	"context"
	"log"

	_ "github.com/lib/pq"
	"github.com/phuc-create/go-simple-crud/server"
)

func main() {
	ctx := context.Background()
	db, err := server.DBConnection()
	if err != nil {
		log.Fatal(err)

	}
	server.New(ctx, db)
}
