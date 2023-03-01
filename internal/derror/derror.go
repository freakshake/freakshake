package derror

import (
	"errors"
	"net/http"
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

var errHTTPStatusMap = map[error]int{
	// Status internal server error.
	ErrInternal: http.StatusInternalServerError,

	// Status unauthorized.
	ErrUnauthorized:       http.StatusUnauthorized,
	ErrInvalidCredentials: http.StatusUnauthorized,

	// Status forbidden.
	ErrInaccessibility: http.StatusForbidden,

	// Status conflict.
	ErrUsernameExists: http.StatusConflict,
	ErrUserExists:     http.StatusConflict,

	// Status bad request.
	ErrInvalidRequest:   http.StatusBadRequest,
	ErrInvalidFirstName: http.StatusBadRequest,
	ErrInvalidLastName:  http.StatusBadRequest,
	ErrInvalidFileID:    http.StatusBadRequest,
	ErrInvalidTime:      http.StatusBadRequest,
	ErrInvalidOffset:    http.StatusBadRequest,
	ErrInvalidLimit:     http.StatusBadRequest,

	// Status not found.
	ErrUnknownFile: http.StatusNotFound,
	ErrUnknownUser: http.StatusNotFound,
}
