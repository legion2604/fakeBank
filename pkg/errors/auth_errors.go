package errors

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrPasswordIncorrect = errors.New("password is incorrect")
	ErrEmailExists       = errors.New("email already exists")
)
