package users

import "errors"

var (
	CouldNotFindAnyUser   = errors.New("Could not find any user!")
	ErrUserAlreadyExist   = errors.New("Users already exist")
	ErrMissingInformation = errors.New("Missing information")
)
