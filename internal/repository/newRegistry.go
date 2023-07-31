package repository

import (
	"database/sql"
	users2 "github.com/phuc-create/go-simple-crud/internal/repository/users"
)

type Registry interface {
	Users() users2.Repository
}

type implement struct {
	db    *sql.DB
	users users2.Repository
}

func (r implement) Users() users2.Repository {
	return r.users
}

func New(db *sql.DB) Registry {
	return implement{db: db, users: users2.New(db)}
}
