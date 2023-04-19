package handlers

import (
	"github.com/phuc-create/go-simple-crud/controllers/user"
)

type UserInput struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Handler struct {
	userServices user.Controllers
}

func New(userSvc user.Controllers) Handler {
	return Handler{
		userServices: userSvc,
	}
}
