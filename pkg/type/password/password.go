package password

import (
	"golang.org/x/crypto/bcrypt"
)

type Password string

func (p Password) Validate() error {
	if len(p) < 8 {
		return ErrInvalidPassword
	}
	return nil
}

func (p Password) Hash() ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
}

func (p Password) CompareWithHashPassword(hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(p)) == nil
}
