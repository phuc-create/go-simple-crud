package users

import (
	"github.com/phuc-create/go-simple-crud/internal/repository/orm"
	"github.com/phuc-create/go-simple-crud/models"
)

func toUser(u *orm.UserAccount) models.User {
	urs := models.User{
		ID:        u.ID,
		Username:  u.Username.String,
		Password:  u.Password.String,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	return urs
}
