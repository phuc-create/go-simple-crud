package users

import (
	"context"

	"github.com/phuc-create/go-simple-crud/internal/repository/orm"
	"github.com/phuc-create/go-simple-crud/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (i implement) GetAllUsers(ctx context.Context) ([]models.User, error) {
	boil.DebugMode = true

	users, err := orm.Users().All(ctx, i.db)
	if err != nil {
		return nil, err
	}

	result := make([]models.User, 0)

	for _, c := range users {
		result = append(result, toUser(c))
	}

	return result, nil
}
