package message_error

import "errors"

var (
	ErrDB               = errors.New("database error")
	ErrCantCreateUserDB = errors.New("can't create user: database error")
	ErrUserAlreadyUsed  = errors.New("can't create user: username has already been used")
	ErrUserInvalid      = errors.New("user is invalid")
	ErrChatInvalid      = errors.New("chat name is invalid")
	ErrCountUsers       = errors.New("two or more users required")
	ErrChatIdInvalid    = errors.New("invalid chat ID")
	ErrUserIdInvalid    = errors.New("invalid user ID")
	ErrUserOrChat       = errors.New("user or chat is invalid")
)
