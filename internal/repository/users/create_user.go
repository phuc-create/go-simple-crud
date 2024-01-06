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
	if users, _ := orm.Users(qm.Where("username=?", user.Username)).All(ctx, i.db); len(users) > 0 {
		return models.User{}, UserAlreadyExist
	}
	var u = orm.User{
		ID:       user.ID,
		Email:    null.StringFrom(user.Email),
		Address:  null.StringFrom(user.Address),
		Phone:    null.StringFrom(user.Phone),
		Country:  null.StringFrom(user.Country),
		Company:  null.StringFrom(user.Company),
		Username: null.StringFrom(user.Username),
		Password: null.StringFrom(user.Username),
	}

	err := u.Insert(ctx, i.db, boil.Infer())
	if err != nil {
		return models.User{}, err
	}
	return toUser(&u), nil
}
