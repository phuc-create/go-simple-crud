package users

import (
	"context"
	"database/sql"
	"errors"
	"github.com/phuc-create/go-simple-crud/internal/repository/orm"
	"github.com/phuc-create/go-simple-crud/models"
	"github.com/phuc-create/go-simple-crud/pkgErrors"
)

func (i implement) GetAllUsers(ctx context.Context) ([]models.User, error) {
	_, err := orm.UserAccounts().All(ctx, i.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []models.User{}, pkgErrors.ErrCouldNotFindAnyUser
		}
	}
	var result []models.User
	//for _, c := range users {
	//	result = append(result, c)
	//}
	return result, nil
}
