package router

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type Router struct {
	router *chi.Mux
	db     *sql.DB
}

func (r Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	r.router.ServeHTTP(writer, request)
}

func NewRouter(db *sql.DB) (*Router, error) {

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	r := &Router{
		router: router,
		db:     db,
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})

	return r, nil
}
