package user

import "errors"

var (
	ErrUserDoesNotExist          = errors.New("user does not exist")
	ErrUserAlreadyExist          = errors.New("user already exist")
	ErrMissingInformation        = errors.New("missing information. please check again")
	ErrPasswordContainWhiteSpace = errors.New("password should not contain white space")
)
