package mobile

import "regexp"

var mobileRegex = regexp.MustCompile(`^(0)?(\d{10})$`)

type MobileNumber string

func (m MobileNumber) Validate() error {
	if !mobileRegex.MatchString(string(m)) {
		return ErrInvalidMobileNumber
	}
	return nil
}
