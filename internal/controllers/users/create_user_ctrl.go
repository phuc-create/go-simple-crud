package users

import (
	"context"
	"github.com/phuc-create/go-simple-crud/models"
)

func (i implement) CreateUserController(ctx context.Context, user models.User) (models.User, error) {
	users, err := i.repo.Users().CreateUser(ctx, user)
	if err != nil {
		return models.User{}, err
	}
	return users, nil
}
