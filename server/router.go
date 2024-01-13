package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	prodCtrls "github.com/phuc-create/go-simple-crud/internal/controllers/products"
	userCtrls "github.com/phuc-create/go-simple-crud/internal/controllers/users"
	"github.com/phuc-create/go-simple-crud/internal/handlers/products"
	"github.com/phuc-create/go-simple-crud/internal/handlers/users"
	"github.com/phuc-create/go-simple-crud/internal/repository"
)

// MasterRouter masterRoute
type MasterRouter struct {
	context        context.Context
	Router         *chi.Mux
	db             *sql.DB
	userHandler    users.UserHandlers
	productHandler products.ProductHandlers
}

type controllers struct {
	usersController    userCtrls.UserControllers
	productsController prodCtrls.ProductControllers
}

func initControllers(db *sql.DB) controllers {
	repo := repository.New(db)
	return controllers{
		usersController:    userCtrls.New(repo),
		productsController: prodCtrls.New(repo),
	}
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

	ctrls := initControllers(db)
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	r := &MasterRouter{
		context:        ctx,
		Router:         router,
		db:             db,
		userHandler:    users.New(ctrls.usersController),
		productHandler: products.New(ctrls.productsController),
	}

	r.initRoutes()
	fmt.Println("server listening & serve at localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func (mr MasterRouter) initRoutes() {
	mr.initUserRoutes()
	mr.initProductRoutes()
}

func (mr MasterRouter) initUserRoutes() {
	pattern := "/api/users"

	mr.Router.Group(func(r chi.Router) {
		r.NotFound(NotFound)
		r.MethodNotAllowed(MethodNotAllowed)
		r.Get(pattern, mr.userHandler.GetAllUser(mr.context))
		r.Post(pattern, mr.userHandler.CreateUser(mr.context))
		//r.Get(pattern+"/{userID}", mr.Handler.GetUserByID)
		//r.Delete(pattern+"/{userID}", mr.Handler.DeleteUser)
		//r.Put(pattern+"/{userID}", mr.Handler.UpdateUserByID)
	})
}
func (mr MasterRouter) initProductRoutes() {
	pattern := "/api/products"

	mr.Router.Group(func(r chi.Router) {
		r.NotFound(NotFound)
		r.MethodNotAllowed(MethodNotAllowed)
		r.Get(pattern, mr.productHandler.GetAllProducts(mr.context))
	})
}
