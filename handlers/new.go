package handlers

import (
	"github.com/phuc-create/go-simple-crud/controllers/user"
)

type Handler struct {
	userServices user.Controllers
}

func New(userSvc user.Controllers) Handler {
	return Handler{
		userServices: userSvc,
	}
}
