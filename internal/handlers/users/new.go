package users

import (
	"github.com/phuc-create/go-simple-crud/internal/controllers/users"
)

type UserHandlers struct {
	controllers users.UserControllers
}

func New(controllers users.UserControllers) UserHandlers {
	return UserHandlers{
		controllers: controllers,
	}
}
