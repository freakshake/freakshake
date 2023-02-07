package derror

import "errors"

var (
	ErrUnknownFile = errors.New("unknown file")
	ErrUnknownUser = errors.New("unknown user")
)

var (
	ErrInvalidRequest     = errors.New("invalid request")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidFirstName   = errors.New("invalid first name")
	ErrInvalidLastName    = errors.New("invalid last name")
	ErrInvalidFileID      = errors.New("invalid file id")
	ErrInvalidTime        = errors.New("invalid time")
	ErrInvalidOffset      = errors.New("invalid offset")
	ErrInvalidLimit       = errors.New("invalid limit")
)

var (
	ErrUnexpected      = errors.New("unexpected error")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrInaccessibility = errors.New("inaccessibility error")
	ErrUsernameExists  = errors.New("username already exists")
	ErrInternal        = errors.New("internal server error")
)
