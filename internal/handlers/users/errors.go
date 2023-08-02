package users

import (
	"errors"
	"fmt"
	"github.com/phuc-create/go-simple-crud/helpers"
	userRepo "github.com/phuc-create/go-simple-crud/internal/repository/users"
	"net/http"
)

var (
	ValidationFail            = "validation_failed"
	CouldNotFindAnyUser       = errors.New("could not find any user")
	MissingInformation        = errors.New("missing information")
	PasswordContainWhiteSpace = errors.New("password should not contain white space")
	AtLeastFiveChars          = errors.New("username should have at least 5 chars")
	MaximumTenChars           = errors.New("username should have maximum 10 chars")
)

var (
	ErrMissingInformation        = &helpers.ErrResponse{Status: http.StatusBadRequest, Code: ValidationFail, Message: "Missing information"}
	ErrCouldNotFindAnyUser       = &helpers.ErrResponse{Status: http.StatusBadRequest, Code: ValidationFail, Message: "Could not find any user!"}
	ErrPasswordContainWhiteSpace = &helpers.ErrResponse{Status: http.StatusBadRequest, Code: ValidationFail, Message: "Password should not contain white space"}
	ErrUserAlreadyExist          = &helpers.ErrResponse{Status: http.StatusBadRequest, Code: ValidationFail, Message: "Users already exist"}
	ErrAtLeastFiveChars          = &helpers.ErrResponse{Status: http.StatusBadRequest, Code: ValidationFail, Message: "Username should have at least 5 chars"}
	ErrMaximumTenChars           = &helpers.ErrResponse{Status: http.StatusBadRequest, Code: ValidationFail, Message: "Username should have maximum 10 chars"}
)

func errorHandler(err error) error {
	fmt.Println(err)
	switch err {
	case CouldNotFindAnyUser:
		return ErrCouldNotFindAnyUser
	case userRepo.UserAlreadyExist:
		return ErrUserAlreadyExist
	case MissingInformation:
		return ErrMissingInformation
	case PasswordContainWhiteSpace:
		return ErrPasswordContainWhiteSpace
	case AtLeastFiveChars:
		return ErrAtLeastFiveChars
	case MaximumTenChars:
		return ErrMaximumTenChars
	}
	return helpers.DefaultError
}
