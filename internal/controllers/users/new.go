package users

import (
	"context"

	"github.com/phuc-create/go-simple-crud/internal/repository"
	"github.com/phuc-create/go-simple-crud/models"
)

type UserControllers interface {
	GetAllUsers(context.Context) ([]models.User, error)
	//GetUserByID(userID string) (models.User, error)
	CreateUser(context.Context, models.User) (models.User, error)
	//DeleteUser(userID string) (bool, error)
	//UpdateUserByID(users *models.User) (models.User, error)
}

type implement struct {
	repo repository.Registry
}

// New Dependency injection
func New(repo repository.Registry) UserControllers {
	return implement{repo: repo}
}
