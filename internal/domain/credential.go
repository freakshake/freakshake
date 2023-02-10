package model

import (
	"github.com/mehdieidi/storm/pkg/type/email"
	"github.com/mehdieidi/storm/pkg/type/mobile"
	"github.com/mehdieidi/storm/pkg/type/password"
)

type Credential struct {
	FirstName      string              `json:"first_name"`
	LastName       string              `json:"last_name"`
	Email          email.Email         `json:"email"`
	MobileNumber   mobile.MobileNumber `json:"mobile_number"`
	Password       password.Password   `json:"-"`
	HashedPassword string              `json:"-"`
}
