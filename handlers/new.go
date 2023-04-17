package handlers

import (
	"github.com/phuc-create/go-simple-crud/controllers"
)

type UserInput struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Handler struct {
	userServices controllers.Controllers
}

func New(userSvc controllers.Controllers) Handler {
	return Handler{
		userServices: userSvc,
	}
}
