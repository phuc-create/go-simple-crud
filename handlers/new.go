package handlers

import (
	"github.com/phuc-create/go-simple-crud/controllers"
)

type Handler struct {
	userServices controllers.Services
}

func New(userSvc controllers.Services) Handler {
	return Handler{
		userServices: userSvc,
	}
}
