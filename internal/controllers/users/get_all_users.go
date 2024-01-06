package users

import (
	"context"

	"github.com/phuc-create/go-simple-crud/models"
)

func (i implement) GetAllUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	users, err := i.repo.Users().GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
