package message_error

import "errors"

var (
	ErrCantCreateUserDB = errors.New("can't create user: database error")
	ErrUserAlreadyUsed  = errors.New("can't create user: username has already been used")
	ErrUserInvalid      = errors.New("can't create user: username must contain from 1 to 50 characters")
)
