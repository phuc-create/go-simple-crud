package users

import (
	"github.com/phuc-create/go-simple-crud/internal/repository/orm"
	"github.com/phuc-create/go-simple-crud/models"
)

func toUser(u *orm.User) models.User {
	urs := models.User{
		ID:        u.ID,
		Username:  u.Username.String,
		Password:  u.Password.String,
		Email:     u.Email.String,
		Address:   u.Address.String,
		Phone:     u.Phone.String,
		Country:   u.Country.String,
		Company:   u.Company.String,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	return urs
}
