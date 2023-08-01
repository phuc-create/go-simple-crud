package users

import (
	"github.com/phuc-create/go-simple-crud/helpers"
	"net/http"
)

const (
	ValidationFail               = "validation_failed"
	CouldNotFindAnyUser          = "Could not find any user!"
	ErrUserAlreadyExist          = "users already exist"
	ErrMissingInformation        = "missing information. please check again"
	ErrPasswordContainWhiteSpace = "password should not contain white space"
)

var (
	ErrCouldNotFindAnyUser = &helpers.ErrResponse{Status: http.StatusBadRequest, Code: ValidationFail, Message: CouldNotFindAnyUser}
)
