package router

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/phuc-create/go-simple-crud/handlers"
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
	prefix := "/api"
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	r := &Router{
		router: router,
		db:     db,
	}

	router.Get(prefix+"/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})
	router.Get(prefix+"/users", handlers.GetUsers)
	router.Post(prefix+"/user", handlers.CreateUser)

	return r, nil
}
