package password

import (
	"golang.org/x/crypto/bcrypt"
)

const minLength = 8

type Password string

func (p Password) Validate() error {
	if len(p) < minLength {
		return ErrInvalidPassword
	}
	return nil
}

func (p Password) Hash() ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
}

func (p Password) CompareWithHashedPassword(hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(p)) == nil
}
