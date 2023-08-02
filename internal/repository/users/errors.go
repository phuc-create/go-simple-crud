package users

import "errors"

var (
	UserAlreadyExist = errors.New("users already exist")
)
