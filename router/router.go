package router

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/phuc-create/go-simple-crud/internal/handlers/users"
	"github.com/phuc-create/go-simple-crud/internal/repository"
	users2 "github.com/phuc-create/go-simple-crud/internal/repository/controllers/users"
	"log"
	"net/http"
)

// MasterRouter masterRoute
type MasterRouter struct {
	context          context.Context
	Router           *chi.Mux
	db               *sql.DB
	usersControllers users2.Controllers
	usersHandler     users.Handler
}
type controllers struct {
	usersController users2.Controllers
}

func initControllers(repo repository.Registry) controllers {
	return controllers{usersController: users2.New(repo)}
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

func New(ctx context.Context, db *sql.DB) {
	repo := repository.New(db)

	controllers := initControllers(repo)
	userHandler := users.New(controllers.usersController)
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	r := &MasterRouter{
		context:          ctx,
		Router:           router,
		db:               db,
		usersHandler:     userHandler,
		usersControllers: controllers.usersController,
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
		r.Get(pattern, mr.usersHandler.GetAllUser(mr.context))
		//r.Post(pattern, mr.Handler.CreateUser)
		//r.Get(pattern+"/{userID}", mr.Handler.GetUserByID)
		//r.Delete(pattern+"/{userID}", mr.Handler.DeleteUser)
		//r.Put(pattern+"/{userID}", mr.Handler.UpdateUserByID)
	})
}
