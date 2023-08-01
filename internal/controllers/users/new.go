package users

import (
	"context"

	"github.com/phuc-create/go-simple-crud/internal/repository"
	"github.com/phuc-create/go-simple-crud/models"
)

type Controllers interface {
	GetAllUsersController(context.Context) ([]models.User, error)
	//GetUserByID(userID string) (models.User, error)
	CreateUserController(context.Context, models.User) (models.User, error)
	//DeleteUser(userID string) (bool, error)
	//UpdateUserByID(users *models.User) (models.User, error)
}

type implement struct {
	repo repository.Registry
}

// New Dependency injection
func New(repo repository.Registry) Controllers {
	return implement{repo: repo}
}
