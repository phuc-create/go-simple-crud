package users

import (
	"errors"
	"fmt"
	"github.com/phuc-create/go-simple-crud/helpers"
	"net/http"
)

var (
	ValidationFail            = "validation_failed"
	CouldNotFindAnyUser       = errors.New("could not find any user")
	ErrUserAlreadyExist       = "Users already exist"
	MissingInformation        = errors.New("Missing information")
	PasswordContainWhiteSpace = errors.New("Password should not contain white space")
)

var (
	ErrMissingInformation        = &helpers.ErrResponse{Status: http.StatusBadRequest, Code: ValidationFail, Message: "Missing information"}
	ErrCouldNotFindAnyUser       = &helpers.ErrResponse{Status: http.StatusBadRequest, Code: ValidationFail, Message: "Could not find any user!"}
	ErrPasswordContainWhiteSpace = &helpers.ErrResponse{Status: http.StatusBadRequest, Code: ValidationFail, Message: "Password should not contain white space"}
)

func errorHandler(err error) error {
	fmt.Println(err)
	switch err {
	case CouldNotFindAnyUser:
		return ErrCouldNotFindAnyUser
	case MissingInformation:
		return ErrMissingInformation
	case PasswordContainWhiteSpace:
		return ErrPasswordContainWhiteSpace
	}
	return helpers.DefaultError
}
