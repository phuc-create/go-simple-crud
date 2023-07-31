package users

import (
	"context"
	"fmt"

	"github.com/friendsofgo/errors"
	"github.com/phuc-create/go-simple-crud/internal/repository/orm"
	"github.com/phuc-create/go-simple-crud/models"
	"github.com/phuc-create/go-simple-crud/pkgErrors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (i implement) GetAllUsers(ctx context.Context) ([]models.User, error) {
	boil.DebugMode = true

	users, err := orm.UserAccounts().All(ctx, i.db)
	fmt.Println(users, err)
	if err != nil {
		return nil, err
	}
	if len(users) < 1 {
		return nil, errors.WithStack(pkgErrors.ErrCouldNotFindAnyUser)
	}

	var result []models.User

	for _, c := range users {
		result = append(result, toUser(c))
	}
	return result, nil
}
