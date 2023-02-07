package gender

import "errors"

var (
	ErrInvalidGender = errors.New("invalid gender")
	ErrUnknownGender = errors.New("unknown gender")
)
