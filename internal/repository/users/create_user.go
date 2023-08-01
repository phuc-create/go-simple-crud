package users

import (
	"context"
	"github.com/phuc-create/go-simple-crud/internal/repository/orm"
	"github.com/phuc-create/go-simple-crud/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (i implement) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	if users, _ := orm.UserAccounts(qm.Where("id=?", user.ID)).All(ctx, i.db); len(users) > 0 {
		return models.User{}, ErrUserAlreadyExist
	}
	var u = orm.UserAccount{
		ID:       user.ID,
		Username: null.String{String: user.Username},
		Password: null.String{String: user.Password},
	}

	err := u.Insert(ctx, i.db, boil.Infer())
	if err != nil {
		return models.User{}, err
	}
	return toUser(&u), nil

}
