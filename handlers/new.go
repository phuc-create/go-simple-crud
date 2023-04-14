package handlers

import (
	"github.com/phuc-create/go-simple-crud/controllers"
)

type UserInput struct {
	Username string
	Password string
}

type Handler struct {
	userServices controllers.Controllers
}

func New(userSvc controllers.Controllers) Handler {
	return Handler{
		userServices: userSvc,
	}
}
