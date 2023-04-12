package controllers

import (
	"database/sql"
	"github.com/phuc-create/go-simple-crud/models"
)

type Services interface {
	GetAllUser() ([]*models.User, error)
	GetUserByID(userID string) (models.User, error)
	CreateUser(user *models.User) (models.User, error)
	DeleteUser(userID string) (bool, error)
	UpdateUser(user *models.User) (models.User, error)
}

type implement struct {
	db *sql.DB
}

// New Dependency injection
func New(db *sql.DB) Services {
	return implement{db: db}
}
