package products

import "database/sql"

type Repository interface {
	GetAllProducts() error
}

func New(db *sql.DB) Repository {
	return implement{
		db: db,
	}
}

type implement struct {
	db *sql.DB
}
