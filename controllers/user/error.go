package user

import "errors"

var (
	// ErrUserDoesNotExist mean the user already exist in DB
	ErrUserDoesNotExist = errors.New("user does not exist")

	// ErrUserAlreadyExist mean the user already exist in DB
	ErrUserAlreadyExist = errors.New("user already exist")

	// ErrMissingInformation mean the information user given missing
	ErrMissingInformation = errors.New("missing information. please check again")

	// ErrPasswordContainWhiteSpace mean the password should not contain white space
	ErrPasswordContainWhiteSpace = errors.New("password should not contain white space")
)
