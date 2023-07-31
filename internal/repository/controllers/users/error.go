package users

import "errors"

var (
	ErrUserDoesNotExist          = errors.New("users does not exist")
	ErrUserAlreadyExist          = errors.New("users already exist")
	ErrMissingInformation        = errors.New("missing information. please check again")
	ErrPasswordContainWhiteSpace = errors.New("password should not contain white space")
)
