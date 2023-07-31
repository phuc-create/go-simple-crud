package users

import (
	"context"
	"database/sql"
	"github.com/phuc-create/go-simple-crud/models"
)

type Repository interface {
	GetAllUsers(ctx context.Context) ([]models.User, error)
}

func New(db *sql.DB) Repository {
	return implement{
		db: db,
	}
}

type implement struct {
	db *sql.DB
}
