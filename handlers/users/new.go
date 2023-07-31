package users

import (
	"github.com/phuc-create/go-simple-crud/internal/repository/controllers/users"
)

type Handler struct {
	controllers users.Controllers
}

func New(controllers users.Controllers) Handler {
	return Handler{
		controllers: controllers,
	}
}
