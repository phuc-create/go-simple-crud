package router

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/phuc-create/go-simple-crud/controllers/user"
	"github.com/phuc-create/go-simple-crud/handlers/users"
	"log"
	"net/http"
)

// MasterRouter masterRoute
type MasterRouter struct {
	Router  *chi.Mux
	DB      *sql.DB
	Handler users.Handler
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

func New(db *sql.DB, controllers user.Controllers) {
	newHandler := users.New(controllers)
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	r := &MasterRouter{
		Router:  router,
		DB:      db,
		Handler: newHandler,
	}
	r.initRoutes()
	fmt.Println("server listening & serve at localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func (mr MasterRouter) initRoutes() {
	mr.initUserRoutes()
}

func (mr MasterRouter) initUserRoutes() {
	pattern := "/api/users"

	mr.Router.Group(func(r chi.Router) {
		r.NotFound(NotFound)
		r.MethodNotAllowed(MethodNotAllowed)
		r.Get(pattern, mr.Handler.GetAllUser)
		r.Post(pattern, mr.Handler.CreateUser)
		//r.Get(pattern+"/{userID}", mr.Handler.GetUserByID)
		//r.Delete(pattern+"/{userID}", mr.Handler.DeleteUser)
		//r.Put(pattern+"/{userID}", mr.Handler.UpdateUserByID)
	})
}
