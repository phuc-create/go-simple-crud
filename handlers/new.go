package handlers

import (
	"github.com/phuc-create/go-simple-crud/controllers"
)

type Handler struct {
	userServices controllers.Controllers
}

func New(userSvc controllers.Controllers) Handler {
	return Handler{
		userServices: userSvc,
	}
}
