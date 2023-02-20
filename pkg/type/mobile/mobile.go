package mobile

import "regexp"

var mobileRegex = regexp.MustCompile(`^(0)?(\d{10})$`)

type Number string

func (m Number) Validate() error {
	if !mobileRegex.MatchString(string(m)) {
		return ErrInvalidMobileNumber
	}
	return nil
}
