package derror

import (
	"errors"
)

var (
	ErrUnknownFile = errors.New("unknown file")
	ErrUnknownUser = errors.New("unknown user")
)

var (
	ErrInvalidRequest     = errors.New("invalid request")
	ErrInvalidFirstName   = errors.New("invalid first name")
	ErrInvalidLastName    = errors.New("invalid last name")
	ErrInvalidFileID      = errors.New("invalid file id")
	ErrInvalidTime        = errors.New("invalid time")
	ErrInvalidOffset      = errors.New("invalid offset")
	ErrInvalidLimit       = errors.New("invalid limit")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

var (
	ErrUnexpected      = errors.New("unexpected error")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrInaccessibility = errors.New("inaccessibility error")
	ErrUsernameExists  = errors.New("username already exists")
	ErrInternal        = errors.New("internal server error")
	ErrUserExists      = errors.New("user already exists")
)
