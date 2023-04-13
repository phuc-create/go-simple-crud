package router

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/phuc-create/go-simple-crud/controllers"
	"github.com/phuc-create/go-simple-crud/handlers"
	"log"
	"net/http"
)

// MasterRouter masterRoute
type MasterRouter struct {
	Router  *chi.Mux
	DB      *sql.DB
	Handler handlers.Handler
}

func (mr MasterRouter) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	mr.Router.ServeHTTP(writer, request)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(405)
	w.Write([]byte("method is not valid"))
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(405)
	w.Write([]byte("method is not valid"))
}

//func NewRouter(db *sql.DB, handler handlers.Handler) (*MasterRouter, error) {
//	prefix := "/api"
//	router := chi.NewRouter()
//	router.Use(middleware.Logger)
//	r := &MasterRouter{
//		Router: router,
//		DB:     db,
//	}
//	router.NotFound(NotFound)
//	router.MethodNotAllowed(MethodNotAllowed)
//
//	router.Get(prefix+"/", func(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte("Hello world!"))
//	})
//	router.Get(prefix+"/users", GetAllUser)
//	router.Post(prefix+"/user", handlers.CreateUser)
//
//	return r, nil
//}

func New(db *sql.DB, controllers controllers.Controllers) {
	newHandler := handlers.New(controllers)
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	r := &MasterRouter{
		Router:  router,
		DB:      db,
		Handler: newHandler,
	}
	r.initRoutes()
	log.Fatal(http.ListenAndServe(":3000", router))
}

func (mr MasterRouter) initRoutes() {
	mr.initUserRoutes()
}

func (mr MasterRouter) initUserRoutes() {
	prefix := "/api"

	mr.Router.Group(func(r chi.Router) {
		r.Get(prefix+"/users", mr.Handler.GetAllUser)
		r.Post(prefix+"/user", mr.Handler.CreateUser)
	})
}
