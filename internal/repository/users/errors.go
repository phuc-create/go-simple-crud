package users

import "errors"

var (
	ErrUserAlreadyExist          = errors.New("Users already exist")
	ErrMissingInformation        = errors.New("Missing information. please check again")
	ErrPasswordContainWhiteSpace = errors.New("Password should not contain white space")
)
