package repository

import (
	"database/sql"

	"github.com/phuc-create/go-simple-crud/internal/repository/products"
	"github.com/phuc-create/go-simple-crud/internal/repository/users"
)

type Registry interface {
	Users() users.Repository
	Products() products.Repository
}

type implement struct {
	db       *sql.DB
	users    users.Repository
	products products.Repository
}

// Users implements Registry.
func (i implement) Users() users.Repository {
	return i.users
}

// Products implements Registry.
func (i implement) Products() products.Repository {
	return i.products
}

func New(db *sql.DB) Registry {
	return implement{db: db, users: users.New(db), products: products.New(db)}
}
