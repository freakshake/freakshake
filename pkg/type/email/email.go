package email

import (
	"net"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type Email string

func (e Email) Validate() error {
	if len(e) < 3 && len(e) > 254 {
		return ErrInvalidEmail
	}

	if !emailRegex.MatchString(string(e)) {
		return ErrInvalidEmail
	}

	parts := strings.Split(string(e), "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return ErrInvalidEmail
	}

	return nil
}
