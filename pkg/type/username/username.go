package username

import (
	"regexp"
)

var usernameRegex = regexp.MustCompile(`^[a-z0-9_-]{5,26}$`)

type Username string

func (u Username) Validate() error {
	if !usernameRegex.MatchString(string(u)) {
		return ErrInvalidUsername
	}
	return nil
}
