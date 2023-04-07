package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phuc-create/go-simple-crud/apis"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/user/find", apis.FindUser).Methods(GET)
	router.HandleFunc("/api/user/users", apis.GetUsers).Methods(GET)

	fmt.Printf("Start server at port http://localhost:5050\n")
	log.Fatal(http.ListenAndServe(":5050", router))
}
