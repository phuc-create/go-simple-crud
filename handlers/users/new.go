package users

import (
	"github.com/phuc-create/go-simple-crud/controllers/user"
)

type Handler struct {
	users user.Controllers
}

func New(userSvc user.Controllers) Handler {
	return Handler{
		users: userSvc,
	}
}
